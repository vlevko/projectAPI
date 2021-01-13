package web

import (
	"log"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/vlevko/projectAPI/postgres"
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
