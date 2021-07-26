package main

import (
	"fmt"
	"os"
)

func main() {
	prenom := "samuel"
	if prenom == "samuel" {
		fmt.Println("félicitations")
	} else if prenom == "sam" {
		fmt.Println("dommage")
	} else {
		fmt.Println("autre")
	}
	if file, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("fichier introuvable")
	} else {
		fmt.Println(file.Size())
	}
}