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

	if err := createSQLiteDatabase(db); err != nil {
		return nil, err
	}

	log.Println("✅ Database initialized successfully")
	return db, nil
}

func createSQLiteDatabase(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS participants (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		nickname TEXT NOT NULL,
		email TEXT NOT NULL,
		region TEXT NOT NULL,
		project_idea TEXT NOT NULL,
		team_preference BOOLEAN NOT NULL,
		desired_teammate TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_email ON participants(email);
	CREATE INDEX IF NOT EXISTS idx_nickname ON participants(nickname);
	CREATE INDEX IF NOT EXISTS idx_region ON participants(region);
	CREATE INDEX IF NOT EXISTS idx_created_at ON participants(created_at);

	CREATE TABLE IF NOT EXISTS roles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS technologies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS participant_roles (
		participant_id INTEGER NOT NULL,
		role_id INTEGER NOT NULL,
		PRIMARY KEY (participant_id, role_id)
	);

	CREATE TABLE IF NOT EXISTS participant_technologies (
		participant_id INTEGER NOT NULL,
		technology_id INTEGER NOT NULL,
		PRIMARY KEY (participant_id, technology_id)
	);
	`

	_, err := db.Exec(query)
	return err
}
