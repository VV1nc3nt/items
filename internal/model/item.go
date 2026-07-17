package model

import "time"

type Item struct {
	ID          int64     `db:"id"`
	Category    string    `db:"category"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ImageKey    string    `db:"image_key"`
	Price       int64     `db:"price"`
	Quantity    int32     `db:"quantity"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ItemCreateInput struct {
	Category    string `db:"category"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ImageKey    string `db:"image_key"`
	Price       int64  `db:"price"`
	Quantity    int32  `db:"quantity"`
}

type ItemGetInput struct {
	ID int64 `db:"id"`
}
