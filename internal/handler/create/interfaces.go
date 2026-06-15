package create

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
)

//mockery:generate: true
type itemManagerService interface {
	Create(ctx context.Context, req *model.ItemInput) (*model.Item, error)
}
