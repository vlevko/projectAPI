package main

// App struct
type App struct {
	Projects []project
}

type project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Columns     []column `json:"columns"`
}

type column struct {
	Name  string `json:"name"`
	Tasks []task `json:"tasks"`
}

type task struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Comments    []string `json:"comments"`
}

type byName []project
