// Package web defines supported routes and implements their handlers
package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vlevko/projectAPI/postgres"
)

// Handler struct holds an HTTP router and a DB store
type Handler struct {
	*chi.Mux
	Store *postgres.Store
}

// GetPort function returns a listining port from the environment variable or default one
func GetPort() string {
	return ":" + postgres.GetEnv("PORT", port)
}

// GetHandler function returns a handler to manage HTTP request
func GetHandler() *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		Store: postgres.NewStore(),
	}

	h.Use(middleware.Logger, middleware.CleanPath)

	h.Route("/projects", func(r chi.Router) {
		r.Get("/", h.ProjectsList)
		r.Post("/", h.ProjectsCreate)
		r.Get("/{id:[0-9]+}", h.ProjectsRead)
		r.Put("/{id:[0-9]+}", h.ProjectsUpdate)
		r.Delete("/{id:[0-9]+}", h.ProjectsDelete)

		r.Get("/{id:[0-9]+}/tasks", h.ProjectTasksList)
		r.Get("/{id:[0-9]+}/columns", h.ColumnsList)
		r.Post("/{id:[0-9]+}/columns", h.ColumnsCreate)
	})

	h.Route("/columns", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", h.ColumnsRead)
		r.Put("/{id:[0-9]+}", h.ColumnsUpdate)
		r.Delete("/{id:[0-9]+}", h.ColumnsDelete)
		r.Get("/{id:[0-9]+}/{position:[0-9]+}", h.ColumnsPosition)

		r.Get("/{id:[0-9]+}/tasks", h.ColumnTasksList)
		r.Post("/{id:[0-9]+}/tasks", h.TasksCreate)
	})

	h.Route("/tasks", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", h.TasksRead)
		r.Put("/{id:[0-9]+}", h.TasksUpdate)
		r.Delete("/{id:[0-9]+}", h.TasksDelete)
		r.Get("/{id:[0-9]+}/{position:[0-9]+}", h.TasksPosition)
		r.Get("/{id:[0-9]+}/columns/{columnID:[0-9]+}", h.TasksStatus)

		r.Get("/{id:[0-9]+}/comments", h.CommentsList)
		r.Post("/{id:[0-9]+}/comments", h.CommentsCreate)
	})

	h.Route("/comments", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", h.CommentsRead)
		r.Put("/{id:[0-9]+}", h.CommentsUpdate)
		r.Delete("/{id:[0-9]+}", h.CommentsDelete)
	})

	h.NotFound(pageNotFoundResponse)

	return h
}

func response(w http.ResponseWriter, status int, v interface{}) {
	r, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(r)
}

func successResponse(w http.ResponseWriter, status int, str string) {
	response(w, status, map[string]string{responseSuccess: str})
}

func errorResponse(w http.ResponseWriter, status int, str string) {
	response(w, status, map[string]string{responseError: str})
}

func pageNotFoundResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, http.StatusNotFound, pageNotFoundError)
}
