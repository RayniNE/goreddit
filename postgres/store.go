package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/raynine/goreddit"
)

func NewStore(postgresConnectionString string) (*Store, error) {
	db, err := sqlx.Open("postgres", postgresConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error while opening database: %v", err)
	}

	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil
}

type Store struct {
	goreddit.ThreadStore
	goreddit.PostStore
	goreddit.CommentStore
}
