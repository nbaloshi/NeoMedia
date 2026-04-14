package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func GetDB(ctx context.Context, connStr string) (*pgxpool.Pool, error){
	var err error
	pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}
	fmt.Println("Connected to database")
	return pool, nil
}

func CloseDB(ctx context.Context, pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
        fmt.Println("Disconnected from database")
	}
}

func Exec(ctx context.Context, pool *pgxpool.Pool, query string, args ...any) error {
	_, err := pool.Exec(ctx, query, args...)
	return err
}

func QueryRow(ctx context.Context, pool *pgxpool.Pool, query string, args ...any) pgx.Row {
	return pool.QueryRow(ctx, query, args...)
}

func Query(ctx context.Context, pool *pgxpool.Pool, query string, args ...any) (pgx.Rows, error) {
    return pool.Query(ctx, query, args...)
}