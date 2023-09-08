package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type Storage struct {
	mu   *sync.Mutex
	pool *pgxpool.Pool
}

func New(conn string) (*Storage, error) {
	p, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &Storage{pool: p, mu: &sync.Mutex{}}, nil
}
