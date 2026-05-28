package main

import (
	"context"
	"net"
	"os"

	"github.com/VV1nc3nt/items/internal/handler/create"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/repository/postgres"
	"github.com/VV1nc3nt/items/internal/service/item_manager"
	setupdb "github.com/VV1nc3nt/items/internal/setup/db"
	"github.com/VV1nc3nt/items/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	repo := postgres.New(pool)
	srv := item_manager.New(repo)
	h := create.New(srv)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error("failed to listen", "err", err)
		return exitErr
	}

	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, h)
	reflection.Register(s)

	log.Info("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve", "err", err)
		return exitErr
	}

	_ = pool
	return exitOK
}
