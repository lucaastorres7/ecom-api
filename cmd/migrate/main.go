package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lucaastorres7/ecom-api/config"
	"github.com/lucaastorres7/ecom-api/database"
)

// New endpoint for creating migrations
func main() {
		db, err := database.NewPostgresStorage(database.DBConfig{
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Host:     config.Envs.DBHost,
		Port:     config.Envs.DBPort,
		Database: config.Envs.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"pgx",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}