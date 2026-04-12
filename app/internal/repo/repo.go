// Package repo implements the data access layer of the application.
// It handles all database queries, transactions, and data mapping,
// providing a clean interface for the service layer to interact with PostgreSQL.
package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Repo {
	return &Repo{
		db: conn,
	}
}

func (r *Repo) DbPing() error {
	err := r.db.Ping(context.Background())
	return err
}
