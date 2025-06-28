package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewConnection(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{DB: db}, nil
}

func (db *DB) RunMigrations() error {
	migrations := []string{
		createUsersTable,
		createDocumentsTable,
		createSummariesTable,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	name TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`

const createDocumentsTable = `
CREATE TABLE IF NOT EXISTS documents (
	id TEXT PRIMARY KEY,
	user_id TEXT NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	type TEXT NOT NULL,
	size INTEGER NOT NULL,
	uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	processed_at DATETIME,
	FOREIGN KEY (user_id) REFERENCES users(id)
);`

const createSummariesTable = `
CREATE TABLE IF NOT EXISTS summaries (
	id TEXT PRIMARY KEY,
	document_id TEXT NOT NULL,
	content TEXT NOT NULL,
	length TEXT NOT NULL,
	keywords TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (document_id) REFERENCES documents(id)
);`