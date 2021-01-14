package web

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/vlevko/projectAPI/postgres"
)

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

var h = GetHandler()

func TestMain(m *testing.M) {
	code := m.Run()
	clearDB()
	h.Store.ProjectStore.DB.Close()
	os.Exit(code)
}

func TestGetPort(t *testing.T) {
	realPort := GetPort()
	expectedPort := ":" + postgres.GetEnv("PORT", port)
	if expectedPort != realPort {
		t.Errorf("Expected port value '%s'. Got '%s\n'", expectedPort, realPort)
	}

}

func clearDB() {
	if _, err := h.Store.ProjectStore.DB.Exec(clearTables); err != nil {
		log.Fatal(err)
	}

	if _, err := h.Store.ProjectStore.DB.Exec(resetPrimaryKeys); err != nil {
		log.Fatal(err)
	}
}

func testRequest(method, url string, body io.Reader) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, url, body)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

func checkResponseCode(expected, got int, t *testing.T) {
	if expected != got {
		t.Errorf("Expected response code '%d'. Got '%d'\n", expected, got)
	}
}
