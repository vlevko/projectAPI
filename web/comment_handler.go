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

// CommentsList function returns a list of task comments or an error message
func (h *Handler) CommentsList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	c, err := h.Store.CommentsByTask(id)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusOK, c)
}

// CommentsCreate function creates and returns a new task comment or an error message
func (h *Handler) CommentsCreate(w http.ResponseWriter, r *http.Request) {
	var c models.Comment
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
	c.TaskID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.CreateComment(&c); err != nil {
		errorResponse(w, http.StatusInternalServerError, defaultError)
		log.Println(err)
		return
	}
	response(w, http.StatusCreated, c)
}

// CommentsRead function returns a task comment or an error message
func (h *Handler) CommentsRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	c, err := h.Store.Comment(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, commentNotFoundError)
		default:
			errorResponse(w, http.StatusNotFound, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, c)
}

// CommentsUpdate function updates and returns a task comment or an error message
func (h *Handler) CommentsUpdate(w http.ResponseWriter, r *http.Request) {
	var c models.Comment
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
	if err := h.Store.UpdateComment(&c); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, commentNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	response(w, http.StatusOK, c)
}

// CommentsDelete function deletes a task comment and returns a success or an error message
func (h *Handler) CommentsDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, badRequestError)
		log.Println(err)
		return
	}
	if err := h.Store.DeleteComment(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, http.StatusNotFound, commentNotFoundError)
		default:
			errorResponse(w, http.StatusInternalServerError, defaultError)
		}
		log.Println(err)
		return
	}
	successResponse(w, http.StatusOK, commentDeletedResponse)
}
