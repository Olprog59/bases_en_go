package main

import (
	"fmt"
	"testing"
)

func TestSplitLine1(t *testing.T) {
	res := SplitLine("Je suis content de vous former au Go!")

	if res != "Je sui" {
		t.Fatalf("Erreur sur SplitLine. got=%v, want=%v", res, "Je sui")
		fmt.Println("erreur 1")
	}
}
func TestSplitLine2(t *testing.T) {
	res := SplitLine("former au Go!")

	if res != "Je sui" {
		t.Fatalf("Erreur sur SplitLine. got=%v, want=%v", res, "Je sui")
		fmt.Println("erreur 2")
	}
}
