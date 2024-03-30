package main

import (
	"gin-note-app/controllers"
	"gin-note-app/middleware"
	"gin-note-app/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	noteService    service.NoteService        = service.New()
	noteController controllers.NoteController = controllers.New(noteService)
)

func setUpLogger() {
	f, _ := os.Create("app-server.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	server := gin.Default()

	setUpLogger()

	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "note app running with gin"})
	})

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, noteController.GetAll())
	})

	server.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, noteController.Save(ctx))
	})

	server.GET("/:id", func(ctx *gin.Context) {
		ctx.JSON(200, noteController.GetNoteById(ctx))
	})
	server.DELETE("/:id", func(ctx *gin.Context) {
		ctx.JSON(200, noteController.DeleteNote(ctx))
	})

	server.Run(":8080")
}
