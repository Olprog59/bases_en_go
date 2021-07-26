package main

import "fmt"

func main() {
	jour := 3

	switch jour {
	case 0, 1, 2, 3, 4:
		fmt.Println("bon courage")
	case 5, 6:
		fmt.Println("bon week-end")
	default:
		fmt.Println("il y a un petit problème")
	}
}
