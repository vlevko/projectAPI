package main

import "unicode/utf8"

func (a *App) readComments(pID, cID, tID int) []string {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				return a.Projects[pID].Columns[cID].Tasks[cID].Comments
			}
		}
	}
	return []string{}
}

func (a *App) createComment(pID, cID, tID int, text string) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) && utf8.RuneCountInString(text) <= 5000 {
				s := []string{text}
				a.Projects[pID].Columns[cID].Tasks[cID].Comments = append(s, a.Projects[pID].Columns[cID].Tasks[cID].Comments...)
			}
		}
	}
}

func (a *App) readComment(pID, cID, tID, id int) string {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				if id >= 0 && id < len(a.Projects[pID].Columns[cID].Tasks[cID].Comments) {
					return a.Projects[pID].Columns[cID].Tasks[cID].Comments[id]
				}
			}
		}
	}
	return ""
}

func (a *App) updateComment(pID, cID, tID, id int, text string) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				if id >= 0 && id < len(a.Projects[pID].Columns[cID].Tasks[cID].Comments) && utf8.RuneCountInString(text) <= 5000 {
					a.Projects[pID].Columns[cID].Tasks[cID].Comments[id] = text
				}
			}
		}
	}
}

func (a *App) deleteComment(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.Projects) {
		if cID >= 0 && cID < len(a.Projects[pID].Columns) {
			if tID >= 0 && tID < len(a.Projects[pID].Columns[cID].Tasks) {
				if id >= 0 && id < len(a.Projects[pID].Columns[cID].Tasks[cID].Comments) {
					a.Projects[pID].Columns[cID].Tasks[tID].Comments = append(a.Projects[pID].Columns[cID].Tasks[tID].Comments[:id], a.Projects[pID].Columns[cID].Tasks[tID].Comments[id+1:]...)
				}
			}
		}
	}
}
