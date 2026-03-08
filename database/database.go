package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	migrate(db)
	return db
}

func migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		title      TEXT NOT NULL,
		completed  INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}
}
