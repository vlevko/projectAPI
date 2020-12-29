package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Run func
func (a *App) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/projects", a.ReadProjects).Methods("GET")
	r.HandleFunc("/projects", a.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id:[0-9]+}", a.ReadProject).Methods("GET")
	r.HandleFunc("/projects/{id:[0-9]+}", a.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id:[0-9]+}", a.DeleteProject).Methods("DELETE")
	r.HandleFunc("/projects/{id:[0-9]+}/columns", a.ReadColumns).Methods("GET")
	r.HandleFunc("/projects/{id:[0-9]+}/columns", a.CreateColumn).Methods("POST")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}", a.ReadColumn).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}", a.UpdateColumn).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}", a.DeleteColumn).Methods("DELETE")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/{id:[0-9]+}", a.ChangeColumnPosition).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/tasks", a.ReadProjectTasks).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks", a.ReadColumnTasks).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks", a.CreateTask).Methods("POST")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}", a.ReadTask).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}", a.UpdateTask).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}", a.DeleteTask).Methods("DELETE")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/{id:[0-9]+}", a.ChangeTaskPosition).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/{id:[0-9]+}/tasks/{tID:[0-9]+}", a.ChangeTaskStatus).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/comments", a.ReadComments).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/comments", a.ReadComments).Methods("POST")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/comments/{id:[0-9]+}", a.ReadComment).Methods("GET")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/comments/{id:[0-9]+}", a.UpdateComment).Methods("PUT")
	r.HandleFunc("/projects/{pID:[0-9]+}/columns/{cID:[0-9]+}/tasks/{tID:[0-9]+}/comments/{id:[0-9]+}", a.UpdateComment).Methods("DELETE")
	r.NotFoundHandler = http.HandlerFunc(sendNotFoundError)
	http.ListenAndServe(":8000", r)
}
