package postgres

import (
	"context"
	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresDB(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

//go:embed queries/up.sql
var upQuery string

func CreateTables(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, upQuery)
	return err
}
