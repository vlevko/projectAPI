package main

import "unicode/utf8"

func (a *App) readProjectTasks(pID int) []task {
	t := []task{}
	if pID >= 0 && pID < len(a.Projects) {
		for cID := range a.Projects[pID].Columns {
			t = append(t, a.readColumnTasks(pID, cID)...)
		}
	}
	return t
}

func (a *App) readColumnTasks(pID, cID int) []task {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			return a.Projects[pID].Columns[cID].Tasks
		}
	}
	return []task{}
}

func (a *App) createTask(pID, cID int, name string, description ...string) {
	if pID >= 0 && pID < len(a.Projects) && utf8.RuneCountInString(name) <= 500 {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			t := task{}
			t.Name = name
			if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 5000 {
				t.Description = description[0]
			}
			a.Projects[pID].Columns[cID].Tasks = append(a.Projects[pID].Columns[cID].Tasks, t)
		}
	}
}

func (a *App) readTask(pID, cID, tID int) task {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				return a.Projects[pID].Columns[cID].Tasks[tID]
			}
		}
	}
	return task{}
}

func (a *App) updateTask(pID, cID, tID int, name string, description ...string) {
	if pID >= 0 && pID < len(a.Projects) && utf8.RuneCountInString(name) <= 500 {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				a.Projects[pID].Columns[cID].Tasks[tID].Name = name
				if len(description) > 0 && utf8.RuneCountInString(description[0]) <= 5000 {
					a.Projects[pID].Columns[cID].Tasks[tID].Description = description[0]
				}
			}
		}
	}
}

func (a *App) deleteTask(pID, cID, tID int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				a.Projects[pID].Columns[cID].Tasks = append(a.Projects[pID].Columns[cID].Tasks[:tID], a.Projects[pID].Columns[cID].Tasks[tID+1:]...)
			}
		}
	}
}

func (a *App) changeTaskPosition(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID != id && tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) && id >= 0 && id < len(a.Projects[pID].Columns[cID].Tasks) {
				t := a.Projects[pID].Columns[cID].Tasks[tID]
				if tID < id {
					copy(a.Projects[pID].Columns[cID].Tasks[tID:id], a.Projects[pID].Columns[cID].Tasks[tID+1:id+1])
					a.Projects[pID].Columns[cID].Tasks[id] = t
				} else {
					copy(a.Projects[pID].Columns[cID].Tasks[id+1:tID+1], a.Projects[pID].Columns[cID].Tasks[id:tID])
					a.Projects[pID].Columns[cID].Tasks[id] = t
				}
			}
		}
	}
}

func (a *App) changeTaskStatus(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID != id && cID >= 0 && cID < len(a.Projects[pID].Columns) && id >= 0 && id < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				t := a.Projects[pID].Columns[cID].Tasks[tID]
				copy(a.Projects[pID].Columns[cID].Tasks[tID:], a.Projects[pID].Columns[cID].Tasks[tID+1:])
				a.Projects[pID].Columns[cID].Tasks = a.Projects[pID].Columns[cID].Tasks[:len(a.Projects[pID].Columns[cID].Tasks)-1]
				a.Projects[pID].Columns[id].Tasks = append(a.Projects[pID].Columns[id].Tasks, t)
			}
		}
	}
}
