package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const table = `CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	surname TEXT
);`

func CreateDB() (*sqlx.DB, error) {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", "data.db?_foreign_keys=1")
	if err != nil {
		return db, fmt.Errorf("storage: create db: %w", err)
	}
	if err = db.Ping(); err != nil {
		return db, fmt.Errorf("storage: create db: %w", err)
	}
	_, err = db.Exec(table)
	if err != nil {
		return nil, fmt.Errorf("storage: create tables: %w", err)
	}
	return db, nil
}
