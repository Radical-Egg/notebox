package database

import "database/sql"

type Note struct {
	ID      int
	Title   string
	Content string
}


func CreateNote(db *sql.DB, note Note) error {
	insertNoteSQL := `INSERT INTO notes (title, content) VALUES (?, ?);`
	_, err := db.Exec(insertNoteSQL, note.Title, note.Content)
	return err
}

// TODO 
// CreateNotesTable

// TODO
// EditNote

// TODO
// DeleteNote

// TODO 
// Choas Mode