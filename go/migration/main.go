package main

import (
	"log"
	"os"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	url := os.Getenv("POSTGRESQL_URL")
	m, err := migrate.New(
		"file://db/migrations",
		url)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}