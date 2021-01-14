package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCommentsList(t *testing.T) {
	clearDB()

	w := testRequest("GET", "/tasks/99999999999999999999/comments", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()

	w = testRequest("GET", "/tasks/1/comments", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
	if body := w.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array '[]'. Got '%s'\n", body)
	}
}

func TestCommentsCreate(t *testing.T) {
	clearDB()

	w := testRequest("POST", "/tasks/1/comments", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"text":"New",}`)
	w = testRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"text":"New"}`)
	w = testRequest("POST", "/tasks/99999999999999999999/comments", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()

	jsonStr = []byte(`{"text":"New"}`)
	w = testRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusCreated, w.Code, t)

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

	w := testRequest("GET", "/comments/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("GET", "/comments/1", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()
	createTestComment()

	w = testRequest("GET", "/comments/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestCommentsUpdate(t *testing.T) {
	clearDB()

	w := testRequest("PUT", "/comments/1", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr := []byte(`{"text":"Update",}`)
	w = testRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"text":"Update"}`)
	w = testRequest("PUT", "/comments/99999999999999999999", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	jsonStr = []byte(`{"text":"Update"}`)
	w = testRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()
	createTestComment()

	jsonStr = []byte(`{"text":"Update"}`)
	w = testRequest("PUT", "/comments/1", bytes.NewBuffer(jsonStr))
	checkResponseCode(http.StatusOK, w.Code, t)
}

func TestCommentsDelete(t *testing.T) {
	clearDB()

	w := testRequest("DELETE", "/comments/99999999999999999999", nil)
	checkResponseCode(http.StatusBadRequest, w.Code, t)

	w = testRequest("DELETE", "/comments/1", nil)
	checkResponseCode(http.StatusNotFound, w.Code, t)

	createTestProject()
	createTestDefaultColumn()
	createTestNewTask()
	createTestComment()

	w = testRequest("DELETE", "/comments/1", nil)
	checkResponseCode(http.StatusOK, w.Code, t)
}

func createTestComment() {
	h.Store.ProjectStore.DB.Exec(`INSERT INTO comments(text, task_id) VALUES('New', 1)`)
}
