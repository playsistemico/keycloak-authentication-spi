package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	createTableSQL := `
    CREATE TABLE users (
        id TEXT NOT NULL PRIMARY KEY,
        name TEXT NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return db
}
