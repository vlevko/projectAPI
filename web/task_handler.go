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

// ProjectTasksList function returns a list of project tasks or an error message
func (h *Handler) ProjectTasksList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	t, err := h.Store.TasksByProject(id)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusOK, t)
}

// ColumnTasksList function returns a list of column tasks or an error message
func (h *Handler) ColumnTasksList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	t, err := h.Store.TasksByColumn(id)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusOK, t)
}

// TasksCreate function creates and returns a new column task or an error message
func (h *Handler) TasksCreate(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	var err error
	t.ColumnID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.CreateTask(&t); err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusCreated, t)
}

// TasksRead function returns a column task or an error message
func (h *Handler) TasksRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	t, err := h.Store.Task(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, taskNotFoundError)
		default:
			errorResponse(w, http.StatusNotFound, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, t)
}

// TasksUpdate function updates and returns a column task or an error message
func (h *Handler) TasksUpdate(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	defer r.Body.Close()
	var err error
	t.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.UpdateTask(&t); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, taskNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, t)
}

// TasksDelete function deletes a column task and returns a success or an error message
func (h *Handler) TasksDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.DeleteTask(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, taskNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, taskDeletedResponse)
}

// TasksPosition function changes a column task position and returns a success or an error message
func (h *Handler) TasksPosition(w http.ResponseWriter, r *http.Request) {
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
	if err := h.Store.ChangeTaskPosition(id, position); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, taskNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, taskPositionResponse)
}

// TasksStatus function changes a task status and returns a success or an error message
func (h *Handler) TasksStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	columnID, err := strconv.Atoi(chi.URLParam(r, "columnID"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.ChangeTaskStatus(id, columnID); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, taskNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, taskStatusResponse)
}
