package postgres

import (
	"fmt"
	"goreddit"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (p *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var t goreddit.Post
	if err := p.Get(&t, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting post : %w", err)
	}
	return t, nil
}
func (p *PostStore) PostsByThread(threadId uuid.UUID) ([]goreddit.Post, error) {
	var tt []goreddit.Post
	if err := p.Select(&tt, `SELECT * FROM posts WHERE thread_id = $1`, threadId); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting post by threadId : %w", err)
	}
	return tt, nil
}

func (s *PostStore) CreatePost(t *goreddit.Post) error {
	if err := s.Get(&t, `INSERT INTO Posts values ($1,$2,$3,$4,$5) RETURNING * `, t.ID, t.ThreadID, t.Title, t.Content, t.Votes); err != nil {
		return fmt.Errorf("error creating Post : %w", err)
	}
	return nil
}
func (s *PostStore) UpdatePost(t *goreddit.Post) error {
	if err := s.Get(&t, `UPDATE Posts SET thread_id=$1, title=$2 ,content=$3, votes=$4 WHERE ID = $5) RETURNING * `, t.ThreadID, t.Title, t.Content, t.Votes, t.ID); err != nil {
		return fmt.Errorf("error updating  Post : %w", err)
	}
	return nil

}
func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE Posts WHERE ID = $1`, id); err != nil {
		return fmt.Errorf("error deleting  Post : %w", err)
	}
	return nil
}
