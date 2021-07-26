package main

import (
	"fmt"
	"strings"
)

func main() {
	display("samuel", "mich", "59", "com", 40, 25)

	nomComplet := format("mich", "sam")
	fmt.Println(nomComplet)

	nom, prenom := splitName1("samuel michaux")
	println(nom, prenom)
	
	nom, _ = splitName1("toto tata")
	println(nom, prenom)
}

func display(nom, prenom, codePostal, ville string, age, num int){
	fmt.Printf("Je m'appelle %s %s. J'habite à %s avec le code postal %s. J'ai %d. Mon num de rue est %d\n",
		nom, prenom, ville, codePostal, age, num)
}

func format(nom, prenom string) (nomComplet string) {
	nomComplet = prenom + " " + nom
	return nomComplet
}

func splitName1(nomComplet string) (prenom, nom string) {
	tab := strings.Split(nomComplet, " ")
	prenom = tab[0]
	nom = tab[1]
	return prenom, nom
}