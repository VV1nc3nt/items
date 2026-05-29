package item_manager

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
)

type Repository interface {
	Create(ctx context.Context, in *model.ItemInput) (model.Item, error)
}
