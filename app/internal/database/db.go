package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConn() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
