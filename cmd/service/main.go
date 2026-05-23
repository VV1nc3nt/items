package main

import (
	"context"
	"os"

	setupdb "github.com/VV1nc3nt/go-marketplace/internal/setup/db"
	"github.com/VV1nc3nt/go-marketplace/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const (
	exitOK  = 0
	exitErr = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	_ = godotenv.Load()

	log := logger.New()
	ctx := context.Background()

	if err := setupdb.Migrate(ctx); err != nil {
		log.Error("migration failed", "err", err)
		return exitErr
	}

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("db connect failed", "err", err)
		return exitErr
	}
	defer pool.Close()

	_ = pool
	return exitOK
}