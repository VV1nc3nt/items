package item_manager

import (
	"context"
	"time"

	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/repository/postgres"
)

type ServiceItemInput struct {
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

type ItemManagerService interface {
	Create(ctx context.Context, req *ServiceItemInput) (*pb.CreateResponse, error)
}

type ItemService struct {
	repo postgres.Repository
}

func New(repo postgres.Repository) *ItemService {
	return &ItemService{repo: repo}
}
