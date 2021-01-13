// Package main runs the projectAPI utility
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/vlevko/projectAPI/postgres"
	"github.com/vlevko/projectAPI/web"
)

func main() {
	postgres.MigrateOrContinue()
	port := web.GetPort()
	handler := web.GetHandler()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func(ch chan os.Signal, DB *sql.DB) {
		<-ch
		DB.Close()
		os.Exit(2)
	}(ch, handler.Store.ProjectStore.DB)
	log.Fatal(http.ListenAndServe(port, handler))
}
