package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestProjectTasksList(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/projects/99999999999999999999/tasks", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	w = testRequest("GET", "/projects/1/tasks", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestColumnTasksList(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/columns/99999999999999999999/tasks", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	w = testRequest("GET", "/columns/1/tasks", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestTasksCreate(t *testing.T) {
	clearDB()

	w := testRequest("POST", "/columns/1/tasks", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"name":"New",}`)
	w = testRequest("POST", "/columns/1/tasks", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"New"}`)
	w = testRequest("POST", "/columns/99999999999999999999/tasks", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	jsonStr = []byte(`{"name":"New"}`)
	w = testRequest("POST", "/columns/1/tasks", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusCreated, w.Code, t)

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

	w := testRequest("GET", "/tasks/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("GET", "/tasks/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()

	w = testRequest("GET", "/tasks/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestTasksUpdate(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/tasks/1", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"name":"Update",}`)
	w = testRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/tasks/99999999999999999999", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestTasksDelete(t *testing.T) {
	clearDB()

	w := testRequest("DELETE", "/tasks/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()

	w = testRequest("DELETE", "/tasks/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestTasksPosition(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/tasks/99999999999999999999/2", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/tasks/1/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/tasks/1/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 2, 1)`)

	w = testRequest("PUT", "/tasks/1/2", nil)
	checkResponseCode(http.StatusOK, w.Code, t)

	w = testRequest("GET", "/columns/1/tasks", nil)
	var m []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m[0]["id"] != 2.0 {
		t.Errorf("Expected first task id to be '2'. Got '%d'\n", m[0]["id"])
	}
	if m[1]["id"] != 1.0 {
		t.Errorf("Expected second task id to be '1'. Got '%d'\n", m[0]["id"])
	}
}

func TestTasksStatus(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/tasks/99999999999999999999/columns/2", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/tasks/1/columns/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/tasks/1/columns/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewColumn()
	createTestNewTask()

	w = testRequest("PUT", "/tasks/1/columns/2", nil)
	checkResponseCode(http.StatusOK, w.Code, t)

	w = testRequest("GET", "/tasks/1", nil)
	var m map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m["columnID"] != 2.0 {
		t.Errorf("Expected task column id to be '2'. Got '%d'\n", m["columnID"])
	}
}

func createTestNewTask() {
	h.Store.ProjectStore.DB.Exec(`INSERT INTO tasks(name, description, position, column_id) VALUES('New', '', 1, 1)`)
}
