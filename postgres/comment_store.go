package postgres

import (
	"database/sql"

	"github.com/vlevko/projectAPI/models"
)

// CommentStore struct holds methods related to the corresponding DB object
type CommentStore struct {
	*sql.DB
}

// Comment function returns the corresponding DB object by its id
func (s *CommentStore) Comment(id int) (models.Comment, error) {
	q :=
		`SELECT text, created_at, task_id
		FROM comments
		WHERE id = $1`
	c := models.Comment{ID: id}
	return c, s.DB.QueryRow(q, c.ID).
		Scan(&c.Text, &c.CreatedAt, &c.TaskID)
}

// CommentsByTask function returns the list of corresponding DB objects
func (s *CommentStore) CommentsByTask(taskID int) ([]models.Comment, error) {
	q :=
		`SELECT id, text, created_at
		FROM comments
		WHERE task_id = $1
		ORDER BY created_at DESC`
	rows, err := s.DB.Query(q, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []models.Comment{}
	for rows.Next() {
		c := models.Comment{TaskID: taskID}
		if err := rows.Scan(&c.ID, &c.Text, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

// CreateComment function creates the new corresponding DB object
func (s *CommentStore) CreateComment(c *models.Comment) error {
	q :=
		`INSERT INTO comments (text, task_id)
		VALUES ($1, $2)
		RETURNING id, created_at`
	return s.DB.QueryRow(q, c.Text, c.TaskID).
		Scan(&c.ID, &c.CreatedAt)
}

// UpdateComment function updates the corresponding DB object
func (s *CommentStore) UpdateComment(c *models.Comment) error {
	q :=
		`UPDATE comments
		SET text = $1
		WHERE id = $2
		RETURNING created_at, task_id`
	return s.DB.QueryRow(q, c.Text, c.ID).
		Scan(&c.CreatedAt, &c.TaskID)
}

// DeleteComment function deletes the corresponding DB object
func (s *CommentStore) DeleteComment(id int) error {
	q :=
		`DELETE FROM comments
		WHERE id = $1`
	r, err := s.DB.Exec(q, id)
	if err == nil {
		i, _ := r.RowsAffected()
		if i == 0 {
			err = sql.ErrNoRows
		}
	}
	return err
}
