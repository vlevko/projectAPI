package postgres

import (
	"database/sql"

	"github.com/vlevko/projectAPI/models"
)

// ColumnStore struct holds methods related to the corresponding DB object
type ColumnStore struct {
	*sql.DB
}

// Column function returns the corresponding DB object by its id
func (s *ColumnStore) Column(id int) (models.Column, error) {
	q :=
		`SELECT name, position, project_id
		FROM columns
		WHERE id = $1`
	c := models.Column{ID: id}
	return c, s.DB.QueryRow(q, c.ID).
		Scan(&c.Name, &c.Position, &c.ProjectID)
}

// ColumnsByProject function returns the list of corresponding DB objects
func (s *ColumnStore) ColumnsByProject(projectID int) ([]models.Column, error) {
	q :=
		`SELECT id, name, position
		FROM columns
		WHERE project_id = $1
		ORDER BY position`
	rows, err := s.DB.Query(q, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := []models.Column{}
	for rows.Next() {
		c := models.Column{ProjectID: projectID}
		if err := rows.Scan(&c.ID, &c.Name, &c.Position); err != nil {
			return nil, err
		}
		columns = append(columns, c)
	}
	if len(columns) == 0 {
		return nil, sql.ErrNoRows
	}
	return columns, nil
}

// CreateColumn function creates the new corresponding DB object
func (s *ColumnStore) CreateColumn(c *models.Column) error {
	q :=
		`INSERT INTO columns (name, position, project_id)
		SELECT $1, COUNT(id) + 1, $2
		FROM columns
		WHERE project_id = $2
		RETURNING id, position`
	return s.DB.QueryRow(q, c.Name, c.ProjectID).
		Scan(&c.ID, &c.Position)
}

// UpdateColumn function updates the corresponding DB object
func (s *ColumnStore) UpdateColumn(c *models.Column) error {
	q :=
		`UPDATE columns
		SET name = $1
		WHERE id = $2
		RETURNING position, project_id`
	return s.DB.QueryRow(q, c.Name, c.ID).
		Scan(&c.Position, &c.ProjectID)
}

// DeleteColumn function deletes the corresponding DB object
func (s *ColumnStore) DeleteColumn(id int) error {
	q :=
		`SELECT *
		FROM delete_column($1)`
	var r int
	err := s.DB.QueryRow(q, id).Scan(&r)
	if err == nil && r == -1 {
		err = sql.ErrNoRows
	}
	return err
}

// ChangeColumnPosition function changes position of the corresponding DB object
func (s *ColumnStore) ChangeColumnPosition(id, position int) error {
	if position < 1 {
		position = 1
	}
	q :=
		`WITH get_pid AS (
			SELECT project_id AS pid
			FROM columns
			WHERE id = $1),
		get_maxpos AS (
			SELECT MAX(position) AS maxpos
			FROM columns
			WHERE project_id = (SELECT pid FROM get_pid)),
		get_pos AS (
			SELECT position AS pos
			FROM columns 
			WHERE id = $1),
		decrease AS (
			UPDATE columns
			SET position = position - 1
			WHERE id <> $1
			AND project_id = (SELECT pid FROM get_pid)
			AND position > (SELECT pos FROM get_pos)
			AND position <= CASE
				WHEN $2 <= (SELECT maxpos FROM get_maxpos)
				THEN $2
				ELSE (SELECT maxpos FROM get_maxpos) END),
		increase AS (
			UPDATE columns
			SET position = position + 1
			WHERE id <> $1
			AND project_id = (SELECT pid FROM get_pid)
			AND position < (SELECT pos FROM get_pos)
			AND position >= CASE
				WHEN $2 <= (SELECT maxpos FROM get_maxpos)
				THEN $2
				ELSE (SELECT maxpos FROM get_maxpos) END)
		UPDATE columns
		SET position = CASE
			WHEN $2 <= (SELECT maxpos FROM get_maxpos)
			THEN $2
			ELSE (SELECT maxpos FROM get_maxpos) END
		WHERE id = $1`
	r, err := s.DB.Exec(q, id, position)
	if err == nil {
		i, _ := r.RowsAffected()
		if i == 0 {
			err = sql.ErrNoRows
		}
	}
	return err
}
