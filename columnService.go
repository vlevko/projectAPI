package main

import "unicode/utf8"

func (a *app) readColumns(pID int) []column {
	if pID >= 0 && pID < len(a.projects) {
		return a.projects[pID].columns
	}
	return []column{}
}

func (a *app) createColumn(pID int, name string) {
	if pID >= 0 && pID < len(a.projects) && utf8.RuneCountInString(name) <= 255 {
		for i := range a.projects[pID].columns {
			if a.projects[pID].columns[i].name == name {
				return
			}
		}
		c := column{}
		c.name = name
		a.projects[pID].columns = append(a.projects[pID].columns, c)
	}
}

func (a *app) readColumn(pID, cID int) column {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			return a.projects[pID].columns[cID]
		}
	}
	return column{}
}

func (a *app) updateColumn(pID, cID int, name string) {
	if pID >= 0 && pID < len(a.projects) && utf8.RuneCountInString(name) <= 255 {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			for i := range a.projects[pID].columns {
				if a.projects[pID].columns[i].name == name {
					return
				}
			}
			a.projects[pID].columns[cID].name = name
		}
	}
}

func (a *app) deleteColumn(pID, cID int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) && len(a.projects[pID].columns) > 1 {
			if len(a.projects[pID].columns[cID].tasks) > 0 {
				id := cID - 1
				if cID == 0 {
					id = len(a.projects[pID].columns) - 1
				}
				a.projects[pID].columns[id].tasks = append(a.projects[pID].columns[id].tasks, a.projects[pID].columns[cID].tasks...)
			}
			a.projects[pID].columns = append(a.projects[pID].columns[:cID], a.projects[pID].columns[cID+1:]...)
		}
	}
}

func (a *app) changeColumnPosition(pID, cID, id int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID != id && cID >= 0 && cID < len(a.projects[pID].columns) && id >= 0 && id < len(a.projects[pID].columns) {
			c := a.projects[pID].columns[cID]
			if cID < id {
				copy(a.projects[pID].columns[cID:id], a.projects[pID].columns[cID+1:id+1])
				a.projects[pID].columns[id] = c
			} else {
				copy(a.projects[pID].columns[id+1:cID+1], a.projects[pID].columns[id:cID])
				a.projects[pID].columns[id] = c
			}
		}
	}
}
