package main

import "unicode/utf8"

func (a *app) readComments(pID, cID, tID int) []string {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				return a.projects[pID].columns[cID].tasks[cID].comments
			}
		}
	}
	return []string{}
}

func (a *app) createComment(pID, cID, tID int, text string) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) && utf8.RuneCountInString(text) <= 5000 {
				s := []string{text}
				a.projects[pID].columns[cID].tasks[cID].comments = append(s, a.projects[pID].columns[cID].tasks[cID].comments...)
			}
		}
	}
}

func (a *app) readComment(pID, cID, tID, id int) string {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				if id >= 0 && id < len(a.projects[pID].columns[cID].tasks[cID].comments) {
					return a.projects[pID].columns[cID].tasks[cID].comments[id]
				}
			}
		}
	}
	return ""
}

func (a *app) updateComment(pID, cID, tID, id int, text string) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				if id >= 0 && id < len(a.projects[pID].columns[cID].tasks[cID].comments) && utf8.RuneCountInString(text) <= 5000 {
					a.projects[pID].columns[cID].tasks[cID].comments[id] = text
				}
			}
		}
	}
}

func (a *app) deleteComment(pID, cID, tID, id int) {
	if pID >= 0 && pID < len(a.projects) {
		if cID >= 0 && cID < len(a.projects[pID].columns) {
			if tID >= 0 && tID < len(a.projects[pID].columns[cID].tasks) {
				if id >= 0 && id < len(a.projects[pID].columns[cID].tasks[cID].comments) {
					a.projects[pID].columns[cID].tasks[tID].comments = append(a.projects[pID].columns[cID].tasks[tID].comments[:id], a.projects[pID].columns[cID].tasks[tID].comments[id+1:]...)
				}
			}
		}
	}
}
