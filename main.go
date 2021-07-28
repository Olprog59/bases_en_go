package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("num", 0, "Test flag int")

	flag.Parse()

	fmt.Printf("%v au carré donne %v\n", *num, *num**num)

	fmt.Println(flag.Args())
}
