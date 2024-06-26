package service

import (
	"book-app/model"
	"fmt"
)

type NoteService interface {
	GetAll() []model.Note
	Save(note model.Note) int
	GetNoteById(id int) (model.Note, error)
	DeleteNote(id int) (int, error)
}

type service struct {
	notes []model.Note
}

func New() NoteService {
	notes := []model.Note{
		model.New("Hello", "Hello world"),
		model.New("Happiness", "Happiness is coding in golang! You should find this happiness too"),
	}
	return &service{notes: notes}
}

func (s *service) GetAll() []model.Note {
	return s.notes
}

func (s *service) Save(note model.Note) int {
	s.notes = append(s.notes, note)

	return len(s.notes)
}

func (s *service) GetNoteById(id int) (model.Note, error) {
	for _, note := range s.notes {
		if note.Id == id {
			return note, nil
		}
	}
	return model.Note{}, fmt.Errorf("note with id: %d not found", id)
}

func (s *service) DeleteNote(id int) (int, error) {
	for index, note := range s.notes {
		if note.Id == id {

			s.notes = append(s.notes[:index], s.notes[index+1:]...)
			return len(s.notes), nil

		}
	}
	return len(s.notes), fmt.Errorf("note with id: %d not found", id)
}
