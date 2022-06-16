package postgres

import (
	"fmt"
	"goreddit"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CommentStore struct {
	*sqlx.DB
}

func (p *CommentStore) Comment(id uuid.UUID) (goreddit.Comment, error) {
	var t goreddit.Comment
	if err := p.Get(&t, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error getting comment : %w", err)
	}
	return t, nil
}
func (p *CommentStore) CommentsByPost(postId uuid.UUID) ([]goreddit.Comment, error) {
	var tt []goreddit.Comment
	if err := p.Select(&tt, `SELECT * FROM comments WHERE post_id = $1`, postId); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error getting post by postId : %w", err)
	}
	return tt, nil
}

func (s *CommentStore) CreateComment(t *goreddit.Comment) error {
	if err := s.Get(&t, `INSERT INTO comments values ($1,$2,$3,$4) RETURNING * `, t.ID, t.PostID, t.Content, t.Votes); err != nil {
		return fmt.Errorf("error creating Comment : %w", err)
	}
	return nil
}
func (s *CommentStore) UpdateComment(t *goreddit.Comment) error {
	if err := s.Get(&t, `UPDATE comments SET post_id=$1, content=$2, votes=$3 WHERE ID = $4) RETURNING * `, t.PostID, t.Content, t.Votes, t.ID); err != nil {
		return fmt.Errorf("error updating  Comment : %w", err)
	}
	return nil

}
func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE comments WHERE ID = $1`, id); err != nil {
		return fmt.Errorf("error deleting  Comment : %w", err)
	}
	return nil
}
