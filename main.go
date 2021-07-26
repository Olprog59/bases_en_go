package main

import "fmt"

func main() {
	golang := []string{"g", "o", "l", "a", "n", "g"}

	fmt.Printf("%v (len='%v', cap='%v')\n", golang, len(golang), cap(golang))

	fmt.Println(golang[0:2])
	fmt.Println(golang[:2])
	fmt.Println(golang[2:])

	tab := make([]string, len(golang))
	copy(tab, golang)

	fmt.Printf("tab: %v (len='%v', cap='%v')\n", tab, len(tab), cap(tab))

	golang[3] = "Sam"
	fmt.Printf("%v (len='%v', cap='%v')\n", golang, len(golang), cap(golang))
	fmt.Printf("tab: %v (len='%v', cap='%v')\n", tab, len(tab), cap(tab))
}
