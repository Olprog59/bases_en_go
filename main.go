package main

import (
	"fmt"
)

func main() {
	var age int // déclaration
	age = 40    // affectation

	// var nom string = "Michaux"

	// var prenom = "Samuel"

	// var prenom, nom string = "samuel", "Michaux"

	var prenom, nom = "samuel", "Michaux"

	codePostal := "59560"

	isOk := false

	fmt.Println(prenom, nom, codePostal, isOk, age)

	var a int8 = -128 // 8 bits => -128 à 127
	var b uint8 = 255 // 8 bits => 0 à 255

	fmt.Println(a, b)
}
