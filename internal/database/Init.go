package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Radical-Egg/notebox/internal/environ"
	_ "github.com/mattn/go-sqlite3"
)

// TODO return database object to reuse

func Init() {
	cfg := environ.LoadConfigs()

	// Open a connection to the SQLite database file
	db, err := sql.Open(cfg.DB_TYPE, cfg.DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO break out create table into its own func

	// Create the "notes" table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			content TEXT
		);`
	_, err = db.Exec(createTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database and table created successfully")
}

