package postgres

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	"github.com/jackc/pgx/v5"
)

func (r *ItemRepository) Create(ctx context.Context, in *model.ItemInput) (model.Item, error) {
	query := `
		INSERT INTO items (category, title, description, image_key, price, quantity)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, category, title, description, image_key, price, quantity, created_at, updated_at
	`

	var item model.Item
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return item, fmt.Errorf("insert item: %w", err)
	}
	defer rows.Close()

	item, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Item])
	if err != nil {
		return item, fmt.Errorf("insert item: %w", err)
	}

	return item, nil
}
