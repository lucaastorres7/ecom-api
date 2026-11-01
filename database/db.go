package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Creates a SQL Storage
func NewPostgresStorage(cfg DBConfig) (*sql.DB, error) {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Path:   cfg.Database,
	}

	// Does not connect to the database
	db, err := sql.Open("pgx", u.String())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
