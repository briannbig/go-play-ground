package services

import (
	"fmt"
	"structs/models"
)

func Greet(u *models.User) {
	fmt.Printf("Hello %s\n", u.UserName)

	printUser(u)
}

func AddFish(u *models.User, fish string) {
	fmt.Printf("Adding %s to preferred fish...\n", fish)
	u.AddFish(fish)

	printUser(u)
}

func printUser(u *models.User) {
	fmt.Println(u)
}
