package item_manager

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
)

func (s *ItemService) Get(ctx context.Context, in *model.ItemGetInput) (*model.Item, error) {

	row, err := s.repo.Get(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("service get item: %w", err)
	}

	return &row, nil
}
