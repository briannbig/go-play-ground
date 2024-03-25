package main

import (
	"structs/models"
	"structs/services"
)

func main() {
	user := &models.User{
		UserName: "Mike",
	}

	services.Greet(user)
}
