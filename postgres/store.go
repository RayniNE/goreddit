package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(postgresConnectionString string) (*Store, error) {
	db, err := sqlx.Open("postgres", postgresConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error while opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return &Store{
		ThreadStore:  &ThreadStore{DB: db},
		PostStore:    &PostStore{db},
		CommentStore: &CommentStore{db},
	}, nil
}

type Store struct {
	*ThreadStore
	*PostStore
	*CommentStore
}
