package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/raynine/goreddit"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) GetThreadById(id uuid.UUID) (goreddit.Thread, error) {
	var thread goreddit.Thread
	if err := s.Get(&thread, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %v", err)
	}

	return thread, nil
}

func (s *ThreadStore) GetAllThreads() ([]goreddit.Thread, error) {
	var threads []goreddit.Thread
	if err := s.Select(&threads, `SELECT * FROM threads`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("error getting threads: %v", err)
	}

	return threads, nil
}

func (s *ThreadStore) CreateThread(thread *goreddit.Thread) error {
	if err := s.Get(thread, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		thread.ID,
		thread.Title,
		thread.Description,
	); err != nil {
		return fmt.Errorf("error while inserting thread: %v", err)
	}

	return nil
}

func (s *ThreadStore) UpdateThread(thread *goreddit.Thread) error {
	if err := s.Get(thread, `UPDATE threads SET title = $2, description = $3 WHERE id = $1 RETURNING *`,
		thread.ID,
		thread.Title,
		thread.Description,
	); err != nil {
		return fmt.Errorf("error while inserting thread: %v", err)
	}

	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error while deleting thread: %v", err)
	}

	return nil
}
