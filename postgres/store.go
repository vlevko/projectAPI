// Package postgres implements the projectAPI DB services
package postgres

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

// Store struct defines the corresponding structures
type Store struct {
	*ProjectStore
	*ColumnStore
	*TaskStore
	*CommentStore
}

// NewStore function creates and returns a pointer to a new store connected to the DB
func NewStore() *Store {
	db, err := sql.Open("postgres", GetEnv("DATABASE_URL", `postgres://postgres:@localhost:5432/postgres?sslmode=disable`))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &Store{
		ProjectStore: &ProjectStore{DB: db},
		ColumnStore:  &ColumnStore{DB: db},
		TaskStore:    &TaskStore{DB: db},
		CommentStore: &CommentStore{DB: db},
	}
}

// GetEnv function returns the value of the given environment variable if it is set or the second argument
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// MigrateOrContinue function applies migrations up or down and exits if corresponding flags set
func MigrateOrContinue() {
	u := flag.Bool("u", false, "migrations 'up'; overrides 'down' if both are set")
	d := flag.Bool("d", false, "migrations 'down'; overridden with 'up' if both are set")
	flag.Parse()
	if *u || *d {
		m, err := migrate.New(
			"file://migrations",
			GetEnv("DATABASE_URL", `postgres://postgres:@localhost:5432/postgres?sslmode=disable`))
		if err != nil {
			log.Fatal(err)
		}
		if *u {
			if err := m.Up(); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := m.Down(); err != nil {
				log.Fatal(err)
			}
		}
		os.Exit(0)
	}
}
