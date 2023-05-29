package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

func CreateNotesTable(db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			content TEXT
		);`
	_, err := db.Exec(createTableSQL)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateNote(db *sql.DB, note Note) error {
	insertNoteSQL := `INSERT INTO notes (title, content) VALUES (?, ?);`
	_, err := db.Exec(insertNoteSQL, note.Title, note.Content)
	return err
}

func GetNoteById(db *sql.DB, id int) (Note, error ) {
	query := "SELECT id, title, content FROM notes where id = ?"

	row := db.QueryRow(query, id)

	note := Note{}

	err := row.Scan(&note.ID, &note.Title, &note.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			return Note{}, fmt.Errorf("Note with ID %d not found", id)
		}
		return Note{}, err
	}

	return note, nil
}

// Function to update a note in the database
func UpdateNote(db *sql.DB, note Note) error {
	query := "UPDATE notes SET title = ?, content = ? WHERE id = ?"

	_, err := db.Exec(query, note.Title, note.Content, note.ID)
	if err != nil {
		return err
	}

	return nil
}

// TODO
// DeleteNote

// TODO 
// Choas Mode