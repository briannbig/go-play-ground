package controllers

import (
	"gin-note-app/model"
	"gin-note-app/service"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	GetAll() []model.Note
	Save(ctx *gin.Context) model.Note
	GetNoteById(ctx *gin.Context) model.Note
	DeleteNote(ctx *gin.Context) int
}

type controller struct {
	service service.NoteService
}

func New(s service.NoteService) NoteController {
	return controller{
		service: s,
	}
}

// DeleteNote implements NoteController.
func (c controller) DeleteNote(ctx *gin.Context) int {
	panic("unimplemented")
}

// GetAll implements NoteController.
func (c controller) GetAll() []model.Note {
	return c.service.GetAll()
}

// GetNoteById implements NoteController.
func (c controller) GetNoteById(ctx *gin.Context) model.Note {
	panic("unimplemented")
}

// Save implements NoteController.
func (c controller) Save(ctx *gin.Context) model.Note {
	var note model.Note
	ctx.BindJSON(&note)

	c.service.Save(note)

	return note
}
