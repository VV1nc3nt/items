package db

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Migrate(ctx context.Context) error {
	// создаем отдельный пул для исполнения миграций и ограничиваем его 30 секундами
	mCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	goose.SetBaseFS(os.DirFS("db/migrations"))
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	return goose.UpContext(mCtx, db, ".")
}
