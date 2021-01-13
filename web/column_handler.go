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

// ColumnsList function returns a list of project columns or an error message
func (h *Handler) ColumnsList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	c, err := h.Store.ColumnsByProject(id)
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
	response(w, http.StatusOK, c)
}

// ColumnsCreate function creates and returns a new project column or an error message
func (h *Handler) ColumnsCreate(w http.ResponseWriter, r *http.Request) {
	var c models.Column
	if r.Body == nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(badRequestError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	var err error
	c.ProjectID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.CreateColumn(&c); err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusCreated, c)
}

// ColumnsRead function returns a project column or an error message
func (h *Handler) ColumnsRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	c, err := h.Store.Column(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, columnNotFoundError)
		default:
			errorResponse(w, http.StatusNotFound, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, c)
}

// ColumnsUpdate function updates and returns a project column or an error message
func (h *Handler) ColumnsUpdate(w http.ResponseWriter, r *http.Request) {
	var c models.Column
	if r.Body == nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(badRequestError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	var err error
	c.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.UpdateColumn(&c); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, columnNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, c)
}

// ColumnsDelete function deletes a project column except the last one moving all project tasks to the one above and returns a success or an error message
func (h *Handler) ColumnsDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.DeleteColumn(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, columnNotDeletedError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, columnDeletedResponse)
}

// ColumnsPosition function changes a project column position and returns a success or an error message
func (h *Handler) ColumnsPosition(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	position, err := strconv.Atoi(chi.URLParam(r, "position"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.ChangeColumnPosition(id, position); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, columnNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, columnPositionResponse)
}
