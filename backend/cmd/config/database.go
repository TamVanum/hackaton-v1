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
		email TEXT NOT NULL,
		region TEXT NOT NULL,
		project_idea TEXT NOT NULL,
		team_preference BOOLEAN NOT NULL,
		desired_teammate TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_email ON registrations(email);
	CREATE INDEX IF NOT EXISTS idx_nickname ON registrations(nickname);
	CREATE INDEX IF NOT EXISTS idx_region ON registrations(region);
	CREATE INDEX IF NOT EXISTS idx_created_at ON registrations(created_at);

	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS technologies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS registration_roles (
		registration_id INTEGER NOT NULL,
		role_id INTEGER NOT NULL,
		PRIMARY KEY (registration_id, role_id)
	);

	CREATE TABLE IF NOT EXISTS registration_technologies (
		registration_id INTEGER NOT NULL,
		technology_id INTEGER NOT NULL,
		PRIMARY KEY (registration_id, technology_id)
	);
	`

	_, err := db.Exec(query)
	return err
}
