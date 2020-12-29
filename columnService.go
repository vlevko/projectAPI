package main

import "unicode/utf8"

func (a *App) readColumns(pID int) []column {
	if pID >= 0 && pID < len(a.Projects) {
		return a.Projects[pID].Columns
	}
	return []column{}
}

func (a *App) createColumn(pID int, name string) {
	if pID >= 0 && pID < len(a.Projects) && utf8.RuneCountInString(name) <= 255 {
		for i := range a.Projects[pID].Columns {
			if a.Projects[pID].Columns[i].Name == name {
				return
			}
		}
		c := column{}
		c.Name = name
		a.Projects[pID].Columns = append(a.Projects[pID].Columns, c)
	}
}

func (a *App) readColumn(pID, cID int) column {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			return a.Projects[pID].Columns[cID]
		}
	}
	return column{}
}

func (a *App) updateColumn(pID, cID int, name string) {
	if pID >= 0 && pID < len(a.Projects) && utf8.RuneCountInString(name) <= 255 {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			for i := range a.Projects[pID].Columns {
				if a.Projects[pID].Columns[i].Name == name {
					return
				}
			}
			a.Projects[pID].Columns[cID].Name = name
		}
	}
}

func (a *App) deleteColumn(pID, cID int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) && len(a.Projects[pID].Columns) > 1 {
			if len(a.Projects[pID].Columns[cID].Tasks) > 0 {
				id := cID - 1
				if cID == 0 {
					id = len(a.Projects[pID].Columns) - 1
				}
				a.Projects[pID].Columns[id].Tasks = append(a.Projects[pID].Columns[id].Tasks, a.Projects[pID].Columns[cID].Tasks...)
			}
			a.Projects[pID].Columns = append(a.Projects[pID].Columns[:cID], a.Projects[pID].Columns[cID+1:]...)
		}
	}
}

func (a *App) changeColumnPosition(pID, cID, id int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID != id && cID >= 0 && cID < len(a.Projects[pID].Columns) && id >= 0 && id < len(a.Projects[pID].Columns) {
			c := a.Projects[pID].Columns[cID]
			if cID < id {
				copy(a.Projects[pID].Columns[cID:id], a.Projects[pID].Columns[cID+1:id+1])
				a.Projects[pID].Columns[id] = c
			} else {
				copy(a.Projects[pID].Columns[id+1:cID+1], a.Projects[pID].Columns[id:cID])
				a.Projects[pID].Columns[id] = c
			}
		}
	}
}
