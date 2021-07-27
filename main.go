package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Email string
}

func (u *User) UpperCase() {
	fmt.Println(strings.ToUpper(u.Name))
	u.Name = strings.ToUpper(u.Name)
}

func main() {
	user := User{"toto", "sam@sam.fr"}

	user.UpperCase()

	fmt.Println(user)
}
