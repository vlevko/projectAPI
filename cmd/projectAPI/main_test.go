package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/vlevko/projectAPI/web"
)

var h = web.GetHandler()

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func clearDB() {
	if _, err := h.Store.ProjectStore.DB.Exec(clearTables); err != nil {
		log.Fatal(err)
	}

	if _, err := h.Store.ProjectStore.DB.Exec(resetPrimaryKeys); err != nil {
		log.Fatal(err)
	}
}

const (
	clearTables = `DELETE FROM comments;
		DELETE FROM tasks;
		DELETE FROM columns;
		DELETE FROM projects;`

	resetPrimaryKeys = `ALTER SEQUENCE comments_id_seq RESTART WITH 1;
		ALTER SEQUENCE tasks_id_seq RESTART WITH 1;
		ALTER SEQUENCE columns_id_seq RESTART WITH 1;
		ALTER SEQUENCE projects_id_seq RESTART WITH 1;`
)

func TestReadProjects(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}

	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestCreateProject(t *testing.T) {
	clearDB()

	jsonStr := []byte(`{"name":"Project", "description":"No. 1"}`)
	r, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusCreated != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusCreated, w.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected project ID to be '1'. Got '%v'\n", m["id"])
	}

	if m["name"] != "Project" {
		t.Errorf("Expected project name to be 'Project'. Got '%v'\n", m["name"])
	}

	if m["description"] != "No. 1" {
		t.Errorf("Expected project description to be 'No. 1'. Got '%v'\n", m["description"])
	}
}

func TestReadProject(t *testing.T) {
	clearDB()

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ := http.NewRequest("GET", "/projects/1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusOK, w.Code)
	}
}

func TestUpdateProject(t *testing.T) {
	clearDB()

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ := http.NewRequest("GET", "/projects/1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	var p map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &p)

	var jsonStr = []byte(`{"name":"Update", "description":"Updated"}`)
	r, _ = http.NewRequest("PUT", "/projects/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusOK, w.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)

	if m["id"] != p["id"] {
		t.Errorf("Expected the id to remain the same '%v'. Got '%v'\n", p["id"], m["id"])
	}

	if m["name"] == p["name"] {
		t.Errorf("Expected the name to change from '%v' to 'Update'. Got '%v'", p["name"], m["name"])
	}

	if m["description"] == p["description"] {
		t.Errorf("Expected the description to change from '%v' to 'Updated'. Got '%v'", p["description"], m["description"])
	}
}

func TestDeleteProject(t *testing.T) {
	clearDB()

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ := http.NewRequest("DELETE", "/projects/1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusOK, w.Code)
	}

	r, _ = http.NewRequest("GET", "/projects/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusNotFound, w.Code)
	}
}
