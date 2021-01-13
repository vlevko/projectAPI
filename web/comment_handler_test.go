package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommentsList(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/tasks/99999999999999999999/comments", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	r, _ = http.NewRequest("GET", "/tasks/1/comments", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestCommentsCreate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("POST", "/tasks/1/comments", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"text":"New",}`)
	r, _ = http.NewRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"text":"New"}`)
	r, _ = http.NewRequest("POST", "/tasks/99999999999999999999/comments", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)

	jsonStr = []byte(`{"text":"New"}`)
	r, _ = http.NewRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusCreated != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusCreated, w.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m["id"] != 1.0 {
		t.Errorf("Expected comment ID to be '1'. Got '%v'\n", m["id"])
	}
	if m["text"] != "New" {
		t.Errorf("Expected comment text to be 'New'. Got '%v'\n", m["text"])
	}
	if m["taskID"] != 1.0 {
		t.Errorf("Expected task ID to be '1'. Got '%v'\n", m["taskID"])
	}
}

func TestCommentsRead(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("GET", "/comments/99999999999999999999", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("GET", "/comments/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO comments(text, task_id) VALUES('New', 1)`)

	r, _ = http.NewRequest("GET", "/comments/1", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestCommentsUpdate(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("PUT", "/comments/1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr := []byte(`{"text":"Update",}`)
	r, _ = http.NewRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"text":"Update"}`)
	r, _ = http.NewRequest("PUT", "/comments/99999999999999999999", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	jsonStr = []byte(`{"text":"Update"}`)
	r, _ = http.NewRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO comments(text, task_id) VALUES('New', 1)`)

	jsonStr = []byte(`{"text":"Update"}`)
	r, _ = http.NewRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}

func TestCommentsDelete(t *testing.T) {
	clearDB()

	r, _ := http.NewRequest("DELETE", "/comments/99999999999999999999", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusBadRequest != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusBadRequest, w.Code)
	}

	r, _ = http.NewRequest("DELETE", "/comments/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusNotFound != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusNotFound, w.Code)
	}

	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)
	h.Store.ProjectStore.DB.Exec(`INSERT INTO comments(text, task_id) VALUES('New', 1)`)

	r, _ = http.NewRequest("DELETE", "/comments/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if http.StatusOK != w.Code {
		t.Errorf("Expected response code '%d'. Got '%d'\n", http.StatusOK, w.Code)
	}
}
