package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectTasksList(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/projects/99999999999999999999/tasks", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	r, _ = http.NewRequest("GET", "/projects/1/tasks", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestColumnTasksList(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/columns/99999999999999999999/tasks", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	r, _ = http.NewRequest("GET", "/columns/1/tasks", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestTasksCreate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("POST", "/columns/1/tasks", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"name":"New",}`)
	r, _ = http.NewRequest("POST", "/columns/1/tasks", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"New"}`)
	r, _ = http.NewRequest("POST", "/columns/99999999999999999999/tasks", bytes.NewBuffer(jsonStr))
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)

	jsonStr = []byte(`{"name":"New"}`)
	r, _ = http.NewRequest("POST", "/columns/1/tasks", bytes.NewBuffer(jsonStr))
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusCreated != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusCreated, w.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m["id"] != 1.0 {
		t.Errorf("Expected task ID to be '1'. Got '%v'\n", m["id"])
	}
	if m["name"] != "New" {
		t.Errorf("Expected task name to be 'New'. Got '%v'\n", m["name"])
	}
	if m["description"] != "" {
		t.Errorf("Expected task description to be ''. Got '%v'\n", m["description"])
	}
	if m["position"] != 1.0 {
		t.Errorf("Expected task position to be '1'. Got '%v'\n", m["position"])
	}
	if m["columnID"] != 1.0 {
		t.Errorf("Expected column ID to be '1'. Got '%v'\n", m["columnID"])
	}
}

func TestTasksRead(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/tasks/99999999999999999999", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("GET", "/tasks/2", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	r, _ = http.NewRequest("GET", "/tasks/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestTasksUpdate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/tasks/1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"name":"Update",}`)
	r, _ = http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/tasks/99999999999999999999", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	jsonStr = []byte(`{"name":"Update"}`)
	r, _ = http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestTasksDelete(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("DELETE", "/tasks/99999999999999999999", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	r, _ = http.NewRequest("DELETE", "/tasks/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestTasksPosition(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/tasks/99999999999999999999/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/tasks/1/99999999999999999999", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/tasks/1/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 2, 1)`)

	r, _ = http.NewRequest("PUT", "/tasks/1/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestTasksStatus(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/tasks/99999999999999999999/columns/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/tasks/1/columns/99999999999999999999", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("PUT", "/tasks/1/columns/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(name, position, project_id) VALUES('New', 2, 1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	r, _ = http.NewRequest("PUT", "/tasks/1/columns/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}
