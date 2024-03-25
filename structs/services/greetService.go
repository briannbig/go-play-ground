package services

import (
	"fmt"
	"structs/models"
)

func Greet(u *models.User) {
	fmt.Printf("Hello %s\n", u.UserName)

	printUser(u)
}

func printUser(u *models.User) {
	fmt.Println(u)
}
