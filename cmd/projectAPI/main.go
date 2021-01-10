// Package main runs the projectAPI utility
package main

import (
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/vlevko/projectAPI/postgres"
	"github.com/vlevko/projectAPI/web"
)

func main() {
	postgres.MigrateOrContinue()
	log.Fatal(
		http.ListenAndServe(
			web.GetPort(),
			web.GetHandler(),
		),
	)
}
