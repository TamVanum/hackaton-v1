package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sql.DB, error) {
	dbPath := "hackathon.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createRegistrationsTable(db); err != nil {
		return nil, err
	}

	log.Println("✅ Database initialized successfully")
	return db, nil
}

func createRegistrationsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		nickname TEXT NOT NULL UNIQUE,
		project_idea TEXT NOT NULL,
		teammate TEXT,
		role TEXT NOT NULL CHECK(role IN ('frontend', 'designer', 'backend', 'fullstack')),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_nickname ON registrations(nickname);
	CREATE INDEX IF NOT EXISTS idx_role ON registrations(role);
	CREATE INDEX IF NOT EXISTS idx_created_at ON registrations(created_at);
	`

	_, err := db.Exec(query)
	return err
}
