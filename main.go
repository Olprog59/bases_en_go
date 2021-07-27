package main

import "fmt"

type User struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	var sam User
	sam.firstname = "samuel"
	//sam.lastname = "michaux"
	sam.age = 40

	fmt.Printf("%#+v\n", sam)

	lony := User{
		"lony",
		"michaux",
		16,
	}
	fmt.Printf("%#+v\n", lony)

	sabrina := User{
		firstname: "Sabrina",
	}
	fmt.Printf("%#+v\n", sabrina)
}
