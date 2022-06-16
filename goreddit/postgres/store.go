package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceConnection string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceConnection)
	if err != nil {
		return nil, fmt.Errorf("Error opening db : %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to db : %w", err)
	}

	return &Store{
		ThreadStore:  &ThreadStore{DB: db},
		PostStore:    &PostStore{DB: db},
		CommentStore: &CommentStore{DB: db},
	}, nil
}

type Store struct {
	*ThreadStore
	*PostStore
	*CommentStore
}
