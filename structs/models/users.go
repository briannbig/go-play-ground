package models

import "fmt"

type User struct {
	UserName string
}

// more like a toString() method in a language like java
func (u *User) String() string {
	return fmt.Sprintf("User{username: %s}", u.UserName)
}
