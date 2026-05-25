package service

import (
	"context"
	"time"

	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/repository"
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

type Service interface {
	Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error)
}

type ItemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}
