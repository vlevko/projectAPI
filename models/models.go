// Package models defines the projectAPI objects and methods
package models

import (
	"time"
)

// Project struct defines the corresponding object
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Column struct defines the corresponding object
type Column struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
	ProjectID int    `json:"projectID"`
}

// Task struct defines the corresponding object
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	ColumnID    int    `json:"columnID"`
}

// Comment struct defines the corresponding object
type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	TaskID    int       `json:"taskID"`
}

// ProjectStore interface defines the corresponding methods
type ProjectStore interface {
	Project(id int) (Project, error)
	Projects() ([]Project, error)
	CreateProject(p *Project) error
	UpdateProject(p *Project) error
	DeleteProject(id int) error
}

// ColumnStore interface defines the corresponding methods
type ColumnStore interface {
	Column(id int) (Column, error)
	ColumnsByProject(projectID int) ([]Column, error)
	CreateColumn(c *Column) error
	UpdateColumn(c *Column) error
	DeleteColumn(id int) error
	ChangeColumnPosition(id, position int) error
}

// TaskStore interface defines the corresponding methods
type TaskStore interface {
	Task(id int) (Task, error)
	TasksByProject(projectID int) ([]Task, error)
	TasksByColumn(columnID int) ([]Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(id int) error
	ChangeTaskPosition(id, position int) error
	ChangeTaskStatus(id, columnID int) error
}

// CommentStore interface defines the corresponding methods
type CommentStore interface {
	Comment(id int) (Comment, error)
	CommentsByTask(taskID int) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id int) error
}
