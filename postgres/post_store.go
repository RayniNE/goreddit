package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/raynine/goreddit"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) GetPostById(id uuid.UUID) (goreddit.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PostStore) GetPostsByThread(id uuid.UUID) ([]goreddit.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PostStore) CreatePost(post *goreddit.Post) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostStore) UpdatePost(post *goreddit.Post) error {
	panic("not implemented") // TODO: Implement
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
