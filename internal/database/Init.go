package database

import (
	"database/sql"

	"github.com/Radical-Egg/notebox/internal/environ"
	_ "github.com/mattn/go-sqlite3"
)


func Connect() (*sql.DB, error) {
	cfg := environ.LoadConfigs()

	// Open a connection to the SQLite database file
	db, err := sql.Open(cfg.DB_TYPE, cfg.DB_NAME)
	if err != nil {
		return nil, err
	}

	return db, err
}

