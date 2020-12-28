package main

type app struct {
	projects []project
}

type project struct {
	name        string
	description string
	columns     []column
}

type column struct {
	name  string
	tasks []task
}

type task struct {
	name        string
	description string
	comments    []string
}

type byName []project
