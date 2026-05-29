package create

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
)

type ItemManagerService interface {
	Create(ctx context.Context, req *model.ItemInput) (*model.Item, error)
}
