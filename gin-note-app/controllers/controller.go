package controllers

import (
	"fmt"
	"gin-note-app/model"
	"gin-note-app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	GetAll() []model.Note
	Save(ctx *gin.Context) model.Note
	GetNoteById(ctx *gin.Context) model.Note
	DeleteNote(ctx *gin.Context) string
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
func (c controller) DeleteNote(ctx *gin.Context) string {
	id_ := ctx.Param("id")

	id, _ := strconv.Atoi(id_)

	res, err := c.service.DeleteNote(id)
	if err != nil {
		return fmt.Sprint(err)
	}

	return fmt.Sprintf("Note deleted, new notes counnt %d", res)
}

// GetAll implements NoteController.
func (c controller) GetAll() []model.Note {
	return c.service.GetAll()
}

// GetNoteById implements NoteController.
func (c controller) GetNoteById(ctx *gin.Context) model.Note {
	id_ := ctx.Param("id")

	id, _ := strconv.Atoi(id_)

	res, err := c.service.GetNoteById(id)
	if err != nil {
		return model.Note{}
	}

	return res

}

// Save implements NoteController.
func (c controller) Save(ctx *gin.Context) model.Note {
	var request requestCreateNote
	ctx.BindJSON(&request)

	note := c.service.Save(
		model.New(request.Title, request.Content),
	)

	return note
}

type requestCreateNote struct {
	Title   string
	Content string
}
