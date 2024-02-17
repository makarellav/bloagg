package server

import (
	"encoding/xml"
	"fmt"
	"github.com/makarellav/bloagg/internal/db"
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/handlers"
	"github.com/makarellav/bloagg/internal/http/middleware"
	"github.com/makarellav/bloagg/internal/services"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Server struct {
	userSrv       *services.UserService
	feedSrv       *services.FeedService
	feedFollowSrv *services.FeedFollowService
	postSrv       *services.PostService
}

func New(userSrv *services.UserService,
	feedSrv *services.FeedService,
	feedFollowSrv *services.FeedFollowService,
	postSrv *services.PostService) *Server {
	return &Server{userSrv: userSrv, feedSrv: feedSrv, feedFollowSrv: feedFollowSrv, postSrv: postSrv}
}

func (s *Server) Run(port string) {
	mux := http.NewServeMux()

	userHandlers := handlers.NewUserHandlers(s.userSrv)
	feedHandlers := handlers.NewFeedHandlers(s.feedSrv, s.feedFollowSrv)
	feedFollowHandlers := handlers.NewFeedFollowHandlers(s.feedFollowSrv)
	postHandlers := handlers.NewPostHandlers(s.postSrv)

	mw := middleware.NewAuthMiddleware(s.userSrv)

	mux.HandleFunc("POST /users", userHandlers.HandleCreate)
	mux.HandleFunc("GET /users", mw.Auth(userHandlers.HandleGetByKey))

	mux.HandleFunc("GET /feeds", feedHandlers.HandleGetAll)
	mux.HandleFunc("POST /feeds", mw.Auth(feedHandlers.HandleCreate))

	mux.HandleFunc("GET /feed_follows", mw.Auth(feedFollowHandlers.HandleGetAllByUserID))
	mux.HandleFunc("POST /feed_follows", mw.Auth(feedFollowHandlers.HandleCreate))
	mux.HandleFunc("DELETE /feed_follows/{feedFollowID}", mw.Auth(feedFollowHandlers.HandleDelete))

	mux.HandleFunc("GET /posts", mw.Auth(postHandlers.HandleGetByUserID))

	log.Fatal(http.ListenAndServe(port, mux))
}

func (s *Server) ScrapeFeeds(limit int, delay time.Duration) {
	ticker := time.NewTicker(delay)

	for ; ; <-ticker.C {
		feeds, err := s.feedSrv.GetNextFeeds(limit)

		if err != nil {
			fmt.Printf("failed to get next feeds %s\n", err)

			continue
		}

		var wg sync.WaitGroup
		for _, feed := range feeds {
			wg.Add(1)
			go s.scrapeFeed(feed, &wg)
		}
		wg.Wait()
	}
}

func (s *Server) scrapeFeed(feed *entity.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	err := s.feedSrv.MarkFeedFetched(feed.ID)

	if err != nil {
		fmt.Printf("failed to mark feed %q fetched: %s\n", feed.Name, err)

		return
	}

	feedData, err := fetchFeed(feed.URL)

	if err != nil {
		fmt.Printf("failed to fetch feed %q: %s\n", feed.Name, err)

		return
	}

	for _, item := range feedData.Channel.Items {
		err = s.postSrv.Create(&item, feed.ID)

		if err != nil {
			if strings.Contains(err.Error(), db.ErrExists) {
				fmt.Printf("found post %q but it already exists, skipping...\n", item.Title)

				continue
			}

			fmt.Printf("failed to save post %q: %s\n", item.Title, err)
		}

		fmt.Printf("found post: %s\n", item.Title)
	}

	fmt.Printf("feed %s collected, found %d posts\n", feed.Name, len(feedData.Channel.Items))
}

func fetchFeed(feedURL string) (*entity.RSSChannel, error) {
	resp, err := http.Get(feedURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var rssChannel entity.RSSChannel

	err = xml.NewDecoder(resp.Body).Decode(&rssChannel)

	if err != nil {
		return nil, err
	}

	return &rssChannel, nil
}
