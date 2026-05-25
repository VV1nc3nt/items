package repository

import (
	"context"
	"fmt"
	"log"
)

func (r *ItemRepository) Create(ctx context.Context, in *CreateItemInput) (Item, error) {
	query := `
		INSERT INTO items (category, title, description, image_key, price, quantity)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, category, title, description, image_key, price, quantity, created_at, updated_at
	`

	var item Item
	if err := r.db.QueryRow(
		ctx,
		query,
		in.Category,
		in.Title,
		in.Description,
		in.ImageKey,
		in.Price,
		in.Quantity,
	).Scan(
		&item.ID,
		&item.Category,
		&item.Title,
		&item.Description,
		&item.ImageKey,
		&item.Price,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	); err != nil {
		return Item{}, fmt.Errorf("repository create item: %w", err)
	}

	log.Printf("created_at: %v, updated_at: %v", item.CreatedAt, item.UpdatedAt)

	return item, nil
}
