package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/makarellav/bloagg/internal/config"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/http/server"
	"github.com/makarellav/bloagg/internal/services"
	"time"
)

func main() {
	ctx := context.TODO()
	cfg, _ := config.Load()
	conn, _ := pgxpool.New(ctx, cfg.DBUrl)

	queries := db.New(conn)

	userSrv := services.NewUserService(queries, ctx)
	feedSrv := services.NewFeedService(queries, ctx)
	feedFollowSrv := services.NewFeedFollowService(queries, ctx)
	postSrv := services.NewPostService(queries, ctx)

	httpServer := server.New(userSrv, feedSrv, feedFollowSrv, postSrv)

	go httpServer.ScrapeFeeds(10, time.Second*10)
	httpServer.Run("localhost:" + cfg.Port)
}
