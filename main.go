package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/makarellav/bloagg/internal/config"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/http/server"
)

func main() {
	ctx := context.TODO()
	cfg, _ := config.Load()
	conn, _ := pgx.Connect(ctx, cfg.DBUrl)
	defer conn.Close(ctx)

	queries := db.New(conn)

	httpServer := server.New(queries)
	httpServer.Run(ctx, "localhost:"+cfg.Port)
}
