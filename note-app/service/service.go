package service

import (
	"book-app/model"
)

type NoteService interface {
	GetAll() []model.Note
	Save(note model.Note) int
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
