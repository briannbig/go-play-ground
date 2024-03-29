package main

import (
	"book-app/controller"
	"book-app/service"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	noteService    service.NoteService       = service.New()
	noteController controller.NoteController = controller.New(noteService)
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to Note App")
	})

	http.HandleFunc("/notes", noteController.GetAll)
	http.HandleFunc("/notes/add", noteController.Save)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Error starting server --- %s", err)
		os.Exit(1)
	} else {
		fmt.Println("Server started")
	}

}
