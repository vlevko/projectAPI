package main

import (
	"sort"
	"unicode/utf8"
)

func (a *app) readProjects() []project {
	return a.projects
}

func (a *app) createProject(name string, description ...string) {
	if utf8.RuneCountInString(name) > 500 {
		return
	}
	p := project{}
	p.name = name
	if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 1000 {
		p.description = description[0]
	}
	c := column{}
	c.name = "Column1"
	p.columns = append(p.columns, c)
	a.projects = append(a.projects, p)
	sort.Sort(byName(a.projects))
}

func (a *app) readProject(id int) project {
	if id >= 0 && id < len(a.projects) {
		return a.projects[id]
	}
	return project{}
}

func (a *app) updateProject(id int, name string, description ...string) {
	if id >= 0 && id < len(a.projects) && utf8.RuneCountInString(name) <= 500 {
		a.projects[id].name = name
		if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 1000 {
			a.projects[id].description = description[0]
		}
		sort.Sort(byName(a.projects))
	}
}

func (a *app) deleteProject(id int) {
	if id >= 0 && id < len(a.projects) {
		a.projects = append(a.projects[:id], a.projects[id+1:]...)
	}
}
