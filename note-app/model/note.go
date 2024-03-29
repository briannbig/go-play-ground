package model

import (
	"encoding/json"
	"log"
	"time"
)

var counter int = 0

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title string, content string) Note {
	counter += 1
	return Note{
		Id:        counter,
		Title:     title,
		Content:   content,
		CreatedAt: time.Time{},
	}
}

func (note Note) String() string {
	out, err := json.MarshalIndent(note, "", "   ")
	if err != nil {
		log.Fatal("Could not marshal data", err)
	}

	return string(out)
}
