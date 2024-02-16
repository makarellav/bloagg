package handlers

import (
	"encoding/json"
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
)

type FeedHandlers struct {
	feedSrv       *services.FeedService
	feedFollowSrv *services.FeedFollowService
}

func NewFeedHandlers(feedSrv *services.FeedService, feedFollowSrv *services.FeedFollowService) *FeedHandlers {
	return &FeedHandlers{feedSrv: feedSrv, feedFollowSrv: feedFollowSrv}
}

func (fh *FeedHandlers) HandleCreate(w http.ResponseWriter, r *http.Request, user *entity.User) {
	var body entity.Feed

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to decode request body")

		return
	}

	body.UserID = user.ID
	newFeed, err := fh.feedSrv.Create(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create a feed")

		return
	}

	newFeedFollow, err := fh.feedFollowSrv.Create(&entity.FeedFollow{
		UserID: user.ID,
		FeedID: newFeed.ID,
	})

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create a feed follow")

		return
	}

	resp := struct {
		Feed       *entity.Feed       `json:"feed"`
		FeedFollow *entity.FeedFollow `json:"feed_follow"`
	}{
		Feed:       newFeed,
		FeedFollow: newFeedFollow,
	}

	response.JSON(w, http.StatusCreated, resp)
}

func (fh *FeedHandlers) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	feeds, err := fh.feedSrv.GetAll()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get feeds")

		return
	}

	response.JSON(w, http.StatusOK, feeds)
}
