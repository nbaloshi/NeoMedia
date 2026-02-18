package sql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func GetDB(ctx context.Context, connStr string) (*pgx.Conn, error){
	var err error
	conn, err = pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}
	fmt.Println("Connected to database")
	return conn, nil
}

func CloseDB(ctx context.Context, conn *pgx.Conn) {
	if conn != nil {
		if err := conn.Close(ctx); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		} else {
			fmt.Println("Disconnected from database")
		}
	}
}

func Exec(ctx context.Context, conn *pgx.Conn, query string, args ...any) error {
	_, err := conn.Exec(ctx, query, args...)
	return err
}

func QueryRow(ctx context.Context, conn *pgx.Conn, query string, args ...any) pgx.Row {
	return conn.QueryRow(ctx, query, args...)
}