package item_manager

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
)

//mockery:generate: true
type repository interface {
	Create(ctx context.Context, in *model.ItemCreateInput) (model.Item, error)
	Get(ctx context.Context, in *model.ItemGetInput) (model.Item, error)
}
