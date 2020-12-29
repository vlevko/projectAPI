package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadComments func
func (a *App) ReadComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	c := a.readComments(pID, cID, tID)
	sendResponse(w, http.StatusOK, c)
}

// CreateComment func
func (a *App) CreateComment(w http.ResponseWriter, r *http.Request) {
	var c struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	a.createComment(pID, cID, tID, c.Text)
	sendResponse(w, http.StatusCreated, "Comment created")
}

// ReadComment func
func (a *App) ReadComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	id, _ := strconv.Atoi(vars["id"])
	c := a.readComment(pID, cID, tID, id)
	sendResponse(w, http.StatusOK, map[string]string{"text": c})
}

// UpdateComment func
func (a *App) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var c struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	id, _ := strconv.Atoi(vars["id"])
	a.updateComment(pID, cID, tID, id, c.Text)
	sendSuccessResponse(w, http.StatusOK, "Comment updated")
}

// DeleteComment func
func (a *App) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := strconv.Atoi(vars["pID"])
	cID, _ := strconv.Atoi(vars["cID"])
	tID, _ := strconv.Atoi(vars["tID"])
	id, _ := strconv.Atoi(vars["id"])
	a.deleteComment(pID, cID, tID, id)
	sendSuccessResponse(w, http.StatusOK, "Comment deleted")
}
