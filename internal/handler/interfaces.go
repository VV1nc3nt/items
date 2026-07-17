package handler

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
)

//mockery:generate: true
type ItemManagerService interface {
	Create(ctx context.Context, req *model.ItemCreateInput) (*model.Item, error)
	Get(ctx context.Context, req *model.ItemGetInput) (*model.Item, error)
}
