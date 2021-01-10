package postgres

import (
	"database/sql"

	"github.com/vlevko/projectAPI/models"
)

// ProjectStore struct holds methods related to the corresponding DB object
type ProjectStore struct {
	*sql.DB
}

// Project function returns the corresponding DB object by its id
func (s *ProjectStore) Project(id int) (models.Project, error) {
	q :=
		`SELECT id, name, description
		FROM projects
		WHERE id = $1`
	var p models.Project
	return p, s.DB.QueryRow(q, id).
		Scan(&p.ID, &p.Name, &p.Description)
}

// Projects function returns the list of corresponding DB objects
func (s *ProjectStore) Projects() ([]models.Project, error) {
	q :=
		`SELECT id, name, description
		FROM projects
		ORDER BY name`
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := []models.Project{}
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

// CreateProject function creates the new corresponding DB object
func (s *ProjectStore) CreateProject(p *models.Project) error {
	q :=
		`WITH create_project AS (
			INSERT INTO projects (name, description)
			VALUES ($1, $2)
			RETURNING id)
		INSERT INTO columns (project_id)
		SELECT id
		FROM create_project
		RETURNING project_id`
	return s.DB.QueryRow(q, p.Name, p.Description).
		Scan(&p.ID)
}

// UpdateProject function updates the corresponding DB object
func (s *ProjectStore) UpdateProject(p *models.Project) error {
	q :=
		`UPDATE projects
		SET name = $1, description = $2
		WHERE id = $3`
	r, err := s.DB.Exec(q, p.Name, p.Description, p.ID)
	if err == nil {
		i, _ := r.RowsAffected()
		if i == 0 {
			err = sql.ErrNoRows
		}
	}
	return err
}

// DeleteProject function deletes the corresponding DB object
func (s *ProjectStore) DeleteProject(id int) error {
	q :=
		`DELETE FROM projects
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
