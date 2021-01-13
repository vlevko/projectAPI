package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectsList(t *testing.T) {
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

func TestProjectsCreate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("POST", "/projects", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"name":"Project",}`)
	r, _ = http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"Project", "description":"No. 1"}`)
	r, _ = http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
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

	jsonStr = []byte(`{"name":""}`)
	r, _ = http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusInternalServerError != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusInternalServerError, w.Code)
	}
}

func TestProjectsRead(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/projects/99999999999999999999", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ = http.NewRequest("GET", "/projects/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusOK, w.Code)
	}

	r, _ = http.NewRequest("GET", "/projects/2", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusNotFound, w.Code)
	}
}

func TestProjectsUpdate(t *testing.T) {
	clearDB()

	jsonStr := []byte(`{"name":"Update",}`)
	r, _ := http.NewRequest("PUT", "/projects/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/projects/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ = http.NewRequest("GET", "/projects/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	var p map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &p)

	jsonStr = []byte(`{"name":"Update", "description":"Updated"}`)

	r, _ = http.NewRequest("PUT", "/projects/99999999999999999999", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/projects/2", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

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

func TestProjectsDelete(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("DELETE", "/projects/99999999999999999999", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)

	r, _ = http.NewRequest("DELETE", "/projects/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusOK, w.Code)
	}

	r, _ = http.NewRequest("DELETE", "/projects/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d\n'", http.StatusNotFound, w.Code)
	}
}
