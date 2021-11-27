package goreddit

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type ThreadStore interface {
	GetThreadById(uuid.UUID) (Thread, error)
	GetAllThreads() ([]Thread, error)
	CreateThread(*Thread) error
	UpdateThread(*Thread) error
	DeleteThread(uuid.UUID) error
}

type PostStore interface {
	GetPostById(uuid.UUID) (Post, error)
	GetPostsByThread(uuid.UUID) ([]Post, error)
	CreatePost(*Post) error
	UpdatePost(*Post) error
	DeletePost(uuid.UUID) error
}

type CommentStore interface {
	GetCommentById(uuid.UUID) (Comment, error)
	GetCommentByPost(uuid.UUID) ([]Comment, error)
	CreateComment(*Comment) error
	UpdateComment(*Comment) error
	DeleteComment(uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
