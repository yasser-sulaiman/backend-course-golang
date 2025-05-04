package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID       int64  `json:"id"`
	PostID   int64  `json:"post_id"`
	UserID   int64  `json:"user_id"`
	Content  string `json:"content"`
	CreatedT string `json:"created_at"`
	User     User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
	SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id
	FROM comments c
	JOIN users ON c.user_id = users.id
	WHERE post_id = $1
	ORDER BY created_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		comment := Comment{}
		comment.User = User{}
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedT,
			&comment.User.Username,
			&comment.User.ID,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil

}
