package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateItemInput struct {
	Category    string
	Title       string
	Description string
	ImageKey    string
	Price       int64
	Quantity    int32
}

type Item struct {
	ID          int64
	Category    string
	Title       string
	Description string
	ImageKey    string
	Price       int64
	Quantity    int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(ctx context.Context, in *CreateItemInput) error
}

type ItemRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{db: db}
}
