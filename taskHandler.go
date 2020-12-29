package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadProjectTasks func
func (a *App) ReadProjectTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	t := a.readProjectTasks(pID)
	sendResponse(w, http.StatusOK, t)
}

// ReadColumnTasks func
func (a *App) ReadColumnTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	t := a.readColumnTasks(pID, cID)
	sendResponse(w, http.StatusOK, t)
}

// CreateTask func
func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	a.createTask(pID, cID, t.Name, t.Description)
	sendSuccessResponse(w, http.StatusCreated, "Task created")
}

// ReadTask func
func (a *App) ReadTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	t := a.readTask(pID, cID, tID)
	sendResponse(w, http.StatusOK, t)
}

// UpdateTask func
func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var t task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	a.updateTask(pID, cID, tID, t.Name, t.Description)
	sendSuccessResponse(w, http.StatusCreated, "Task updated")
}

// DeleteTask func
func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	a.deleteTask(pID, cID, tID)
	sendSuccessResponse(w, http.StatusOK, "Task deleted")
}

// ChangeTaskPosition func
func (a *App) ChangeTaskPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	id, _ := strconv.Atoi(vars["id"])
	a.changeTaskPosition(pID, cID, tID, id)
	sendSuccessResponse(w, http.StatusOK, "Task position changed")
}

// ChangeTaskStatus func
func (a *App) ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	id, _ := strconv.Atoi(vars["id"])
	a.changeTaskStatus(pID, cID, tID, id)
	sendSuccessResponse(w, http.StatusOK, "Task status changed")
}
