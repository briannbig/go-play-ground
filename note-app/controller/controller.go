package controller

import (
	"book-app/model"
	"book-app/service"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type NoteController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	GetNoteById(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.NoteService
}

func New(service service.NoteService) NoteController {
	return controller{
		service: service,
	}
}

func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	notes := c.service.GetAll()

	notesStr, err := json.MarshalIndent(notes, "", "   ")
	if err != nil {
		log.Println("Could not marshal notes")
	}
	io.WriteString(w, string(notesStr))
}

func (c controller) Save(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Could not read body: ", err)
	}

	fmt.Printf("Recieved payload... %s\n", body)

	request := RequestCreateNote{}

	err = json.Unmarshal(body, &request)

	if err != nil {
		log.Fatalf("Could not unmarshal body: %s\n", err)
	}

	res := c.service.Save(model.New(request.Title, request.Content))

	io.WriteString(w, fmt.Sprintf("New notes size: %d\n", res))
}

func (c controller) GetNoteById(w http.ResponseWriter, r *http.Request) {
	QUERY_STR_ID := "id"

	hasId := r.URL.Query().Has(QUERY_STR_ID)

	if !hasId {
		fmt.Println("Id missing in request params")
		io.WriteString(w, "id not specified in request params")
		return
	}

	id := r.URL.Query().Get(QUERY_STR_ID)

	id_, _ := strconv.Atoi(id)

	note, err := c.service.GetNoteById(id_)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error --- %s", err))
		return
	}

	io.WriteString(w, note.String())

}

// DeleteNote implements NoteController.
func (c controller) DeleteNote(w http.ResponseWriter, r *http.Request) {
	QUERY_STR_ID := "id"

	hasId := r.URL.Query().Has(QUERY_STR_ID)

	if !hasId {
		fmt.Println("Id missing in request params")
		io.WriteString(w, "id not specified in request params")
		return
	}

	id := r.URL.Query().Get(QUERY_STR_ID)

	id_, _ := strconv.Atoi(id)

	res, err := c.service.DeleteNote(id_)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error --- %s", err))
		return
	}

	io.WriteString(w, fmt.Sprintf("New notes size: %d\n", res))
}

type RequestCreateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (requestCreateNote RequestCreateNote) String() string {

	out, err := json.MarshalIndent(requestCreateNote, "", "   ")
	if err != nil {
		log.Println("Could not unmarshal request")
	}

	return string(out)

}
