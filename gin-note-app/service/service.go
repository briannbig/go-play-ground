package service

import (
	"fmt"
	"gin-note-app/data"
	"gin-note-app/model"
)

type NoteService interface {
	GetAll() []model.Note
	Save(note model.Note) model.Note
	GetNoteById(id int) (model.Note, error)
	DeleteNote(id int) (int, error)
}

type service struct {
	notes    []model.Note
	dataBase data.DataBase
}

func New() NoteService {
	notes := []model.Note{
		model.New("Hello", "Hello world"),
		model.New("Happiness", "Happiness is coding in golang! You should find this happiness too"),
	}
	return &service{notes: notes, dataBase: data.New()}
}

func (s *service) GetAll() []model.Note {
	notes := []model.Note{}

	rows, err := s.dataBase.Connection.Query("SELECT * FROM notes")
	if err != nil {
		fmt.Printf("error fetching notes from database --- %s\n", err)
	}

	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			fmt.Printf("Error scanning row --- %s\n", err)
		}

		notes = append(notes, note)
	}
	
	return notes
}

func (s *service) Save(note model.Note) model.Note {

	note_, err := note.Persist(&s.dataBase)
	if err != nil {
		fmt.Printf("Could not persist note: --- %s\n", err)
		return note
	}

	return note_
}

func (s *service) GetNoteById(id int) (model.Note, error) {

	var note model.Note
	err := s.dataBase.Connection.QueryRow("SELECT * FROM notes WHERE id = ?", id).Scan(
		&note.Id, &note.Content, &note.Title)

	if err != nil {
		return note, fmt.Errorf("error fetching note with id: %d --- %s", id, err)
	}

	return note, nil
}

func (s *service) DeleteNote(id int) (int, error) {

	_, err := s.dataBase.Connection.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return -1, fmt.Errorf("error deleting note --- %s", err)
	}

	return 1, nil

}
