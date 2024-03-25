package main

import (
	"structs/models"
	"structs/services"
	"time"
)

func main() {
	user := &models.User{
		UserName:  "Mike",
		Password:  "StrongPassword",
		CreatedAt: time.Now(),
	}

	services.Greet(user)

	services.AddFish(user, "Tilapia")
	services.AddFish(user, "Nile Perch")
}
