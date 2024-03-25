package models

import (
	"encoding/json"
	"log"
	"time"
)

type User struct {
	UserName      string    `json:"name"`
	Password      string    `json:"-"`                       // let serializer ignore this field
	PreferredFish []string  `json:"preferredFish,omitempty"` // let serializer ignore if empty
	CreatedAt     time.Time `json:"createdAt"`
}

// more like a toString() method in a language like java
func (u *User) String() string {
	//return a json representation of the user struct
	out, err := json.MarshalIndent(u, "", "   ")

	if err != nil {
		log.Println(err)
	}
	return string(out)
}

func (u *User) AddFish(fish string) {
	u.PreferredFish = append(u.PreferredFish, fish)
}
