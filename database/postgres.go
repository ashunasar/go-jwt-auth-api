package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  refresh_token TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)

	if err != nil {
		return nil, err
	}

	return db, err

}
