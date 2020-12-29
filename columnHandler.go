package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadColumns func
func (a *App) ReadColumns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	c := a.readColumns(id)
	sendResponse(w, http.StatusOK, c)
}

// CreateColumn func
func (a *App) CreateColumn(w http.ResponseWriter, r *http.Request) {
	var c column
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a.createColumn(id, c.Name)
	sendSuccessResponse(w, http.StatusCreated, "Column created")
}

// ReadColumn func
func (a *App) ReadColumn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	c := a.readColumn(pID, cID)
	sendResponse(w, http.StatusOK, c)
}

// UpdateColumn func
func (a *App) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	var c column
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	a.updateColumn(pID, cID, c.Name)
	sendSuccessResponse(w, http.StatusOK, "Column updated")
}

// DeleteColumn func
func (a *App) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	a.deleteColumn(pID, cID)
	sendSuccessResponse(w, http.StatusOK, "Column deleted")
}

// ChangeColumnPosition func
func (a *App) ChangeColumnPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	id, _ := strconv.Atoi(vars["id"])
	a.changeColumnPosition(pID, cID, id)
	sendSuccessResponse(w, http.StatusOK, "Column position changed")
}
