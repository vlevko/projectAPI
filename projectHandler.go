package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadProjects func
func (a *App) ReadProjects(w http.ResponseWriter, r *http.Request) {
	projects := a.readProjects()
	sendResponse(w, http.StatusOK, projects)
}

// CreateProject func
func (a *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	var p project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	a.createProject(p.Name, p.Description)
	sendSuccessResponse(w, http.StatusCreated, "Project created")
}

// ReadProject func
func (a *App) ReadProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p := a.readProject(id)
	sendResponse(w, http.StatusOK, p)
}

// UpdateProject func
func (a *App) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var p project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a.updateProject(id, p.Name, p.Description)
	sendSuccessResponse(w, http.StatusOK, "Project updated")
}

// DeleteProject func
func (a *App) DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a.deleteProject(id)
	sendSuccessResponse(w, http.StatusOK, "Project deleted")
}
