package server

import (
	"context"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/http/handlers"
	"github.com/makarellav/bloagg/internal/http/middleware"
	"github.com/makarellav/bloagg/internal/services"
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

	userHandlers := handlers.NewUserHandlers(userSrv)
	feedHandlers := handlers.NewFeedHandlers(feedSrv)

	mw := middleware.NewAuthMiddleware(userSrv)

	mux.HandleFunc("POST /users", userHandlers.HandleCreate)
	mux.HandleFunc("GET /users", mw.Auth(userHandlers.HandleGetByKey))

	mux.HandleFunc("POST /feeds", mw.Auth(feedHandlers.HandleCreate))

	http.ListenAndServe(port, mux)
}
