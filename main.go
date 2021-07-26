package main

import "fmt"

func main() {
	fmt.Println("Start main()")
	defer fmt.Println("End main")

	a := true

	if a {
		fmt.Println("je suis dans une condition!!")
		return
	}

}
