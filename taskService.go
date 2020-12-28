package main

import "unicode/utf8"

func (a *app) readProjectTasks(pID int) []task {
	t := []task{}
	if pID >= 0 && pID < len(a.projects) {
		for cID := range a.projects[pID].columns {
			t = append(t, a.readColumnTasks(pID, cID)...)
		}
	}
	return t
}

func (a *app) readColumnTasks(pID, cID int) []task {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			return a.projects[pID].columns[cID].tasks
		}
	}
	return []task{}
}

func (a *app) createTask(pID, cID int, name string, description ...string) {
	if pID >= 0 && pID < len(a.projects) && utf8.RuneCountInString(name) <= 500 {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			t := task{}
			t.name = name
			if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 5000 {
				t.description = description[0]
			}
			a.projects[pID].columns[cID].tasks = append(a.projects[pID].columns[cID].tasks, t)
		}
	}
}

func (a *app) readTask(pID, cID, tID int) task {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				return a.projects[pID].columns[cID].tasks[tID]
			}
		}
	}
	return task{}
}

func (a *app) updateTask(pID, cID, tID int, name string, description ...string) {
	if pID >= 0 && pID < len(a.projects) && utf8.RuneCountInString(name) <= 500 {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				a.projects[pID].columns[cID].tasks[tID].name = name
				if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 5000 {
					a.projects[pID].columns[cID].tasks[tID].description = description[0]
				}
			}
		}
	}
}

func (a *app) deleteTask(pID, cID, tID int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				a.projects[pID].columns[cID].tasks = append(a.projects[pID].columns[cID].tasks[:tID], a.projects[pID].columns[cID].tasks[tID+1:]...)
			}
		}
	}
}

func (a *app) changeTaskPosition(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID != id && tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) && id >= 0 && id < len(a.projects[pID].columns[cID].tasks) {
				t := a.projects[pID].columns[cID].tasks[tID]
				if tID < id {
					copy(a.projects[pID].columns[cID].tasks[tID:id], a.projects[pID].columns[cID].tasks[tID+1:id+1])
					a.projects[pID].columns[cID].tasks[id] = t
				} else {
					copy(a.projects[pID].columns[cID].tasks[id+1:tID+1], a.projects[pID].columns[cID].tasks[id:tID])
					a.projects[pID].columns[cID].tasks[id] = t
				}
			}
		}
	}
}

func (a *app) changeTaskStatus(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID != id && cID >= 0 && cID < len(a.projects[pID].columns) && id >= 0 && id < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				t := a.projects[pID].columns[cID].tasks[tID]
				copy(a.projects[pID].columns[cID].tasks[tID:], a.projects[pID].columns[cID].tasks[tID+1:])
				a.projects[pID].columns[cID].tasks = a.projects[pID].columns[cID].tasks[:len(a.projects[pID].columns[cID].tasks)-1]
				a.projects[pID].columns[id].tasks = append(a.projects[pID].columns[id].tasks, t)
			}
		}
	}
}
