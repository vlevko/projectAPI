package main

import (
	"sort"
	"unicode/utf8"
)

func (a *App) readProjects() []project {
	return a.Projects
}

func (a *App) createProject(name string, description ...string) {
	if utf8.RuneCountInString(name) > 500 {
		return
	}
	p := project{}
	p.Name = name
	if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 1000 {
		p.Description = description[0]
	}
	c := column{}
	c.Name = "Column1"
	p.Columns = append(p.Columns, c)
	a.Projects = append(a.Projects, p)
	sort.Sort(byName(a.Projects))
}

func (a *App) readProject(id int) project {
	if id >= 0 && id < len(a.Projects) {
		return a.Projects[id]
	}
	return project{}
}

func (a *App) updateProject(id int, name string, description ...string) {
	if id >= 0 && id < len(a.Projects) && utf8.RuneCountInString(name) <= 500 {
		a.Projects[id].Name = name
		if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 1000 {
			a.Projects[id].Description = description[0]
		}
		sort.Sort(byName(a.Projects))
	}
}

func (a *App) deleteProject(id int) {
	if id >= 0 && id < len(a.Projects) {
		a.Projects = append(a.Projects[:id], a.Projects[id+1:]...)
	}
}
