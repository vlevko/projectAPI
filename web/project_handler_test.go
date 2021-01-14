package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestProjectsList(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/projects", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestProjectsCreate(t *testing.T) {
	clearDB()

	w := testRequest("POST", "/projects", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"name":"Project",}`)
	w = testRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"Project", "description":"No. 1"}`)
	w = testRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusCreated, w.Code, t)

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
	w = testRequest("POST", "/projects", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusInternalServerError, w.Code, t)
}

func TestProjectsRead(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/projects/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()

	w = testRequest("GET", "/projects/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)

	w = testRequest("GET", "/projects/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)
}

func TestProjectsUpdate(t *testing.T) {
	clearDB()

	jsonStr := []byte(`{"name":"Update",}`)
	w := testRequest("PUT", "/projects/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/projects/1", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()

	w = testRequest("GET", "/projects/1", nil)
	var p map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &p)

	jsonStr = []byte(`{"name":"Update", "description":"Updated"}`)

	w = testRequest("PUT", "/projects/99999999999999999999", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/projects/2", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusNotFound, w.Code, t)

	w = testRequest("PUT", "/projects/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusOK, w.Code, t)

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

	w := testRequest("DELETE", "/projects/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()

	w = testRequest("DELETE", "/projects/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)

	w = testRequest("DELETE", "/projects/1", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)
}

func createTestProject() {
	h.Store.ProjectStore.DB.Exec(`INSERT INTO projects(name, description) VALUES('Project', 'No. 1')`)
}
