package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	"github.com/jackc/pgx/v5"
)

func (r *ItemRepository) Get(ctx context.Context, in *model.ItemGetInput) (model.Item, error) {
	query := `
		SELECT id, category, title, description, image_key, price, quantity, created_at, updated_at
		FROM items 
		WHERE id = $1
	`

	var item model.Item
	rows, err := r.db.Query(ctx, query, in.ID)
	if err != nil {
		return item, fmt.Errorf("db get item: %w", err)
	}
	defer rows.Close()

	item, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Item])
	if errors.Is(err, pgx.ErrNoRows) {
		return item, model.DatabaseError
	}
	if err != nil {
		return item, fmt.Errorf("db get item: %w", err)
	}

	return item, nil
}
