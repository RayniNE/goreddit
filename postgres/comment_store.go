package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/raynine/goreddit"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) GetCommentById(id uuid.UUID) (goreddit.Comment, error) {
	var comment goreddit.Comment
	if err := s.Get(&comment, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error while getting comment: %v", err)
	}

	return comment, nil
}

func (s *CommentStore) GetCommentByPost(id uuid.UUID) ([]goreddit.Comment, error) {
	var comments []goreddit.Comment
	if err := s.Select(&comments, `SELECT * FROM comments WHERE post_id = $1`, id); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error while getting comments: %v", err)
	}

	return comments, nil
}

func (s *CommentStore) CreateComment(comment *goreddit.Comment) error {
	if err := s.Get(comment, `INSERT INTO comments VALUES($1, $2, $3, $4) RETURNING *`,
		comment.ID,
		comment.PostID,
		comment.Content,
		comment.Votes,
	); err != nil {
		return fmt.Errorf("error while inserting comment: %v", err)
	}

	return nil
}

func (s *CommentStore) UpdateComment(comment *goreddit.Comment) error {
	if err := s.Get(comment, `UPDATE comments SET id = $1, post_id = $2, content = $3, votes = $4) RETURNING *`,
		comment.ID,
		comment.PostID,
		comment.Content,
		comment.Votes,
	); err != nil {
		return fmt.Errorf("error while inserting comment: %v", err)
	}

	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error while deleting comment: %v", err)
	}

	return nil
}
