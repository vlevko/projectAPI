package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestColumnsList(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/projects/99999999999999999999/columns", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("GET", "/projects/1/columns", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	r, _ = http.NewRequest("GET", "/projects/1/columns", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}

	if body := w.Body.String(); body != `[{"id":1,"name":"ToDo","position":1,"projectID":1}]` {
		t.Errorf(`Expected an empty array '[{"id":1,"name":"ToDo","position":1,"projectID":1}]'. Got '%s'\n`, body)
	}
}

func TestColumnsCreate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("POST", "/projects/1/columns", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"name":"New",}`)
	r, _ = http.NewRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"New"}`)
	r, _ = http.NewRequest("POST", "/projects/99999999999999999999/columns", bytes.NewBuffer(jsonStr))
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	jsonStr = []byte(`{"name":"ToDo"}`)
	r, _ = http.NewRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusInternalServerError != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusInternalServerError, w.Code)
	}

	jsonStr = []byte(`{"name":"New"}`)
	r, _ = http.NewRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusCreated != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusCreated, w.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m["id"] != 3.0 {
		t.Errorf("Expected column ID to be '3'. Got '%v'\n", m["id"])
	}
	if m["name"] != "New" {
		t.Errorf("Expected column name to be 'New'. Got '%v'\n", m["name"])
	}
	if m["position"] != 2.0 {
		t.Errorf("Expected column position to be '2'. Got '%v'\n", m["position"])
	}
	if m["projectID"] != 1.0 {
		t.Errorf("Expected project ID to be '1'. Got '%v'\n", m["projectID"])
	}
}

func TestColumnsRead(t *testing.T) {
	clearDB()

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	r, _ := http.NewRequest("GET", "/columns/99999999999999999999", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("GET", "/columns/2", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	r, _ = http.NewRequest("GET", "/columns/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestColumnsUpdate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/columns/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"name":"Update",}`)
	r, _ = http.NewRequest("PUT", "/columns/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/columns/99999999999999999999", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/columns/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(name, position, project_id) VALUES('New', 2, 1)`)

	jsonStr = []byte(`{"name":"ToDo"}`)
	r, _ = http.NewRequest("PUT", "/columns/2", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusInternalServerError != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusInternalServerError, w.Code)
	}

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/columns/2", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestColumnsDelete(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("DELETE", "/columns/99999999999999999999", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("DELETE", "/columns/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	r, _ = http.NewRequest("DELETE", "/columns/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(name, position, project_id) VALUES('New', 2, 1)`)

	r, _ = http.NewRequest("DELETE", "/columns/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestColumnsPosition(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/columns/99999999999999999999/2", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/columns/1/99999999999999999999", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/columns/1/2", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(name, position, project_id) VALUES('New', 2, 1)`)

	r, _ = http.NewRequest("PUT", "/columns/1/2", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}
