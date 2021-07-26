package main

import "fmt"

func main() {
	names := []string{
		"Samuel",
		"Nicolas",
		"Patrick",
		"Lony",
		"Perséa",
	}

	//for i := 0; i < len(names); i++ {
	//	fmt.Printf("Name : %s\n", names[i])
	//}

	//for index, value := range names {
	//	fmt.Printf("Name : %s à l'index %d\n", value, index)
	//}

	for _, value := range names {
		fmt.Printf("Name : %s\n", value)
	}
	fmt.Println("----------------------")
	for index := range names {
		fmt.Printf("Name : %s\n", names[index])
	}
}
