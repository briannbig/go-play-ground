package model

import (
	"encoding/json"
	"fmt"
	"gin-note-app/data"
	"log"
)

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(title string, content string) Note {
	return Note{
		Title:   title,
		Content: content,
	}
}

func (note Note) String() string {
	out, err := json.MarshalIndent(note, "", "   ")
	if err != nil {
		log.Fatal("Could not marshal data", err)
	}

	return string(out)
}

func (note Note) Persist(db *data.DataBase) (Note, error) {

	statement, err := db.Connection.Prepare("INSERT INTO notes(title, content) VALUES(?, ?)")
	if err != nil {
		return Note{}, fmt.Errorf("could not prepare sql query --- %s", err)
	}

	res, execErr := statement.Exec(note.Title, note.Content)
	if execErr != nil {
		return Note{}, fmt.Errorf("could not persist note to database --- %s", execErr)
	}

	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		return Note{}, fmt.Errorf("could not get last inserted ID: %s", lastIdErr)
	}

	var savedNote Note
	err = db.Connection.QueryRow("SELECT * FROM notes WHERE id = ?", lastId).
		Scan(&savedNote.Id, &savedNote.Title, &savedNote.Content)
	if err != nil {
		return Note{}, fmt.Errorf("could not fetch saved note from database: %s", err)
	}

	return savedNote, nil

}
