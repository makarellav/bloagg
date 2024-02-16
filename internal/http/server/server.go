package server

import (
	"context"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/http/handlers"
	"github.com/makarellav/bloagg/internal/http/middleware"
	"github.com/makarellav/bloagg/internal/services"
	"log"
	"net/http"
)

type Server struct {
	db *db.Queries
}

func New(db *db.Queries) *Server {
	return &Server{db: db}
}

func (s *Server) Run(ctx context.Context, port string) {
	mux := http.NewServeMux()

	userSrv := services.NewUserService(s.db, ctx)
	feedSrv := services.NewFeedService(s.db, ctx)
	feedFollowSrv := services.NewFeedFollowService(s.db, ctx)

	userHandlers := handlers.NewUserHandlers(userSrv)
	feedHandlers := handlers.NewFeedHandlers(feedSrv, feedFollowSrv)
	feedFollowHandlers := handlers.NewFeedFollowHandlers(feedFollowSrv)

	mw := middleware.NewAuthMiddleware(userSrv)

	mux.HandleFunc("POST /users", userHandlers.HandleCreate)
	mux.HandleFunc("GET /users", mw.Auth(userHandlers.HandleGetByKey))

	mux.HandleFunc("GET /feeds", feedHandlers.HandleGetAll)
	mux.HandleFunc("POST /feeds", mw.Auth(feedHandlers.HandleCreate))

	mux.HandleFunc("GET /feed_follows", mw.Auth(feedFollowHandlers.HandleGetAllByUserID))
	mux.HandleFunc("POST /feed_follows", mw.Auth(feedFollowHandlers.HandleCreate))
	mux.HandleFunc("DELETE /feed_follows/{feedFollowID}", mw.Auth(feedFollowHandlers.HandleDelete))

	log.Fatal(http.ListenAndServe(port, mux))
}
