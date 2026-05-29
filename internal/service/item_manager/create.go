package item_manager

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
)

func (s *ItemService) Create(ctx context.Context, in *model.ItemInput) (*model.Item, error) {

	row, err := s.repo.Create(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("service create item: %w", err)
	}

	return &row, nil
}
