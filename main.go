package main

import "fmt"

func main() {
	// i++ => i = i + 1 => i += 1
	for i := 0; i < 10; i++ {
		fmt.Printf("Index : %d\n", i)
	}

	i := 0
	for i < 4 {
		fmt.Printf("While => index : %d\n", i)
		i++
	}

	i = 0
	for {
		i++
		if i%2 == 1 {
			continue
		}
		fmt.Printf("Loop %d", i)
		if i >= 7 {
			break
		}
	}
}
