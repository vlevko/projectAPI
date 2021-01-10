package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vlevko/projectAPI/models"
)

// ProjectsList function returns a list of projects or an error message
func (h *Handler) ProjectsList(w http.ResponseWriter, r *http.Request) {
	pp, err := h.Store.Projects()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusOK, pp)
}

// ProjectsCreate function creates and returns a new project with default column or an error message
func (h *Handler) ProjectsCreate(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	if err := h.Store.CreateProject(&p); err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusCreated, p)
}

// ProjectsRead function returns a project or an error message
func (h *Handler) ProjectsRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	p, err := h.Store.Project(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, projectNotFoundError)
		default:
			errorResponse(w, http.StatusNotFound, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, p)
}

// ProjectsUpdate function updates and returns a project or an error message
func (h *Handler) ProjectsUpdate(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	var err error
	p.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	err = h.Store.UpdateProject(&p)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, projectNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, p)
}

// ProjectsDelete function deletes a project and returns a success or an error message
func (h *Handler) ProjectsDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.DeleteProject(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, projectNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, projectDeletedResponse)
}
