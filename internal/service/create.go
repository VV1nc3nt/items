package service

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/repository"
)

func (s *ItemService) Create(ctx context.Context, input CreateItemInput) (Item, error) {
	repoInput := repository.CreateItemInput{
		Category:    input.Category,
		Title:       input.Title,
		Description: input.Description,
		ImageKey:    input.ImageKey,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	row, err := s.repo.Create(ctx, &repoInput)
	if err != nil {
		return Item{}, fmt.Errorf("service create item: %w", err)
	}

	item := Item{
		ID:          row.ID,
		Category:    row.Category,
		Title:       row.Title,
		Description: row.Description,
		ImageKey:    row.ImageKey,
		Price:       row.Price,
		Quantity:    row.Quantity,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}

	return item, nil
}
