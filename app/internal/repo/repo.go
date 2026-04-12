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
