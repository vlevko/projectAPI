package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestColumnsList(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/projects/99999999999999999999/columns", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("GET", "/projects/1/columns", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	w = testRequest("GET", "/projects/1/columns", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
	if body := w.Body.String(); body != `[{"id":1,"name":"ToDo","position":1,"projectID":1}]` {
		t.Errorf(`Expected an empty array '[{"id":1,"name":"ToDo","position":1,"projectID":1}]'. Got '%s'\n`, body)
	}
}

func TestColumnsCreate(t *testing.T) {
	clearDB()

	w := testRequest("POST", "/projects/1/columns", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"name":"New",}`)
	w = testRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"New"}`)
	w = testRequest("POST", "/projects/99999999999999999999/columns", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	jsonStr = []byte(`{"name":"ToDo"}`)
	w = testRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusInternalServerError, w.Code, t)

	jsonStr = []byte(`{"name":"New"}`)
	w = testRequest("POST", "/projects/1/columns", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusCreated, w.Code, t)

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

	createTestProject()
	createTestDefaultColumn()

	w := testRequest("GET", "/columns/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("GET", "/columns/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	w = testRequest("GET", "/columns/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestColumnsUpdate(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/columns/1", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"name":"Update",}`)
	w = testRequest("PUT", "/columns/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/columns/99999999999999999999", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/columns/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewColumn()

	jsonStr = []byte(`{"name":"ToDo"}`)
	w = testRequest("PUT", "/columns/2", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusInternalServerError, w.Code, t)

	jsonStr = []byte(`{"name":"Update"}`)
	w = testRequest("PUT", "/columns/2", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestColumnsDelete(t *testing.T) {
	clearDB()

	w := testRequest("DELETE", "/columns/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("DELETE", "/columns/1", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()

	w = testRequest("DELETE", "/columns/1", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestNewColumn()

	w = testRequest("DELETE", "/columns/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestColumnsPosition(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/columns/99999999999999999999/2", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/columns/1/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("PUT", "/columns/1/2", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewColumn()

	w = testRequest("PUT", "/columns/1/2", nil)
	checkResponseCode(http.StatusOK, w.Code, t)

	w = testRequest("GET", "/projects/1/columns", nil)
	var m []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &m)
	if m[0]["id"] != 2.0 {
		t.Errorf("Expected first column id to be '2'. Got '%d'\n", m[0]["id"])
	}
	if m[1]["id"] != 1.0 {
		t.Errorf("Expected second column id to be '1'. Got '%d'\n", m[0]["id"])
	}
}

func createTestDefaultColumn() {
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(project_id) VALUES(1)`)
}

func createTestNewColumn() {
	h.Store.ProjectStore.DB.Exec(`INSERT INTO columns(name, position, project_id) VALUES('New', 2, 1)`)
}
