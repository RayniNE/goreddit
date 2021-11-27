package postgres

import (
	"fmt"

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
	var post goreddit.Post
	if err := s.Get(&post, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error while getting post: %v", err)
	}

	return post, nil
}

func (s *PostStore) GetPostsByThread(id uuid.UUID) ([]goreddit.Post, error) {
	var posts []goreddit.Post
	if err := s.Select(&posts, `SELECT * FROM posts WHERE thread_id = $1`, id); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error while getting posts: %v", err)
	}
	return posts, nil
}

func (s *PostStore) CreatePost(post *goreddit.Post) error {
	if err := s.Get(post, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		post.ID,
		post.ThreadID,
		post.Title,
		post.Content,
		post.Votes,
	); err != nil {
		return fmt.Errorf("error while inserting post: %v", err)
	}

	return nil
}

func (s *PostStore) UpdatePost(post *goreddit.Post) error {
	if err := s.Get(post, `UPADTE posts SET thread_id = $2, title = $3, content = $4, votes = $5 WHERE id = $1 RETURNING *`,
		post.ID,
		post.ThreadID,
		post.Title,
		post.Content,
		post.Votes,
	); err != nil {
		return fmt.Errorf("error while inserting post: %v", err)
	}

	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error while deleting post: %v", err)
	}

	return nil
}
