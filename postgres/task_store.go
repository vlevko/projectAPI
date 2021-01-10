package postgres

import (
	"database/sql"

	"github.com/vlevko/projectAPI/models"
)

// TaskStore struct holds methods related to the corresponding DB object
type TaskStore struct {
	*sql.DB
}

// Task function returns the corresponding DB object by its id
func (s *TaskStore) Task(id int) (models.Task, error) {
	q :=
		`SELECT name, description, position, column_id
		FROM tasks
		WHERE id = $1`
	t := models.Task{ID: id}
	return t, s.DB.QueryRow(q, t.ID).
		Scan(&t.Name, &t.Description, &t.Position, &t.ColumnID)
}

// TasksByProject function returns the list of corresponding DB objects
func (s *TaskStore) TasksByProject(projectID int) ([]models.Task, error) {
	q :=
		`SELECT t.id, t.name, t.description, t.position, t.column_id 
		FROM tasks AS t
		INNER JOIN columns AS c
		ON c.id = t.column_id
		WHERE t.column_id = ANY(
			SELECT c.id 
			FROM columns 
			WHERE project_id = $1)
		ORDER BY c.position, t.position`
	rows, err := s.DB.Query(q, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Position, &t.ColumnID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// TasksByColumn function returns the list of corresponding DB objects
func (s *TaskStore) TasksByColumn(columnID int) ([]models.Task, error) {
	q :=
		`SELECT id, name, description, position 
		FROM tasks 
		WHERE column_id = $1 
		ORDER BY position`
	rows, err := s.DB.Query(q, columnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []models.Task{}
	for rows.Next() {
		t := models.Task{ColumnID: columnID}
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Position); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// CreateTask function creates the new corresponding DB object
func (s *TaskStore) CreateTask(t *models.Task) error {
	q :=
		`INSERT INTO tasks(name, description, position, column_id)
		SELECT $1, $2, COUNT(id) + 1, $3
		FROM tasks
		WHERE column_id = $3
		RETURNING id, position`
	return s.DB.QueryRow(q, t.Name, t.Description, t.ColumnID).
		Scan(&t.ID, &t.Position)
}

// UpdateTask function updates the corresponding DB object
func (s *TaskStore) UpdateTask(t *models.Task) error {
	q :=
		`UPDATE tasks
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING position, column_id`
	return s.DB.QueryRow(q, t.Name, t.Description, t.ID).
		Scan(&t.Position, &t.ColumnID)
}

// DeleteTask function deletes the corresponding DB object
func (s *TaskStore) DeleteTask(id int) error {
	q :=
		`WITH delete_task AS (
			DELETE FROM tasks
			WHERE id = $1
			RETURNING position)
		UPDATE tasks 
		SET position = position - 1
		WHERE position > (SELECT position FROM delete_task)`
	_, err := s.DB.Exec(q, id)
	return err
}

// ChangeTaskPosition function changes position of the corresponding DB object
func (s *TaskStore) ChangeTaskPosition(id, position int) error {
	if position < 1 {
		position = 1
	}
	q :=
		`WITH get_cid AS (
			SELECT column_id AS cid
			FROM tasks 
			WHERE id = $1),
		get_maxpos AS (
			SELECT MAX(position) AS maxpos
			FROM tasks
			WHERE column_id = (SELECT cid FROM get_cid)),
		get_pos AS (
			SELECT position AS pos
			FROM tasks
			WHERE id = $1),
		decrease AS (
			UPDATE tasks
			SET position = position - 1
			WHERE id <> $1
			AND column_id = (SELECT cid FROM get_cid)
			AND position > (SELECT pos FROM get_pos)
			AND position <= CASE
				WHEN $2 <= (SELECT maxpos FROM get_maxpos)
				THEN $2
				ELSE (SELECT maxpos FROM get_maxpos) END),
		increase AS (
			UPDATE tasks
			SET position = position + 1
			WHERE id <> $1
			AND column_id = (SELECT cid FROM get_cid)
			AND position < (SELECT pos FROM get_pos)
			AND position >= CASE
				WHEN $2 <= (SELECT maxpos FROM get_maxpos)
				THEN $2
				ELSE (SELECT maxpos FROM get_maxpos) END)
		UPDATE tasks
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

// ChangeTaskStatus function changes status of the corresponding DB object
func (s *TaskStore) ChangeTaskStatus(id, columnID int) error {
	q :=
		`WITH update_position AS (
			UPDATE tasks
			SET position = position - 1
			WHERE column_id = (
				SELECT column_id
				FROM tasks
				WHERE id = $1)
			AND position > (
				SELECT position
				FROM tasks
				WHERE id = $1))
		UPDATE tasks
		SET position = (
				SELECT COUNT(id) + 1
				FROM tasks
				WHERE column_id = $2),
			column_id = $2
		WHERE id = $1`
	r, err := s.DB.Exec(q, id, columnID)
	if err == nil {
		i, _ := r.RowsAffected()
		if i == 0 {
			err = sql.ErrNoRows
		}
	}
	return err
}
