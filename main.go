package main

import "fmt"

func main() {
	var age = 40

	fmt.Println(age)
	fmt.Println(&age)

	var yearOld = age
	fmt.Println(yearOld)
	fmt.Println(&yearOld)

	yearOld = 25
	fmt.Println(yearOld)
	fmt.Println(&yearOld)
	fmt.Println(age)
	fmt.Println(&age)

	var num1 = 15
	fmt.Printf("num1 = %v\n", num1)
	fmt.Printf("num1 = %v\n", &num1)

	var num2 = &num1
	fmt.Printf("num1 = %v\n", num1)
	fmt.Printf("num1 type = %T\n", num1)
	fmt.Printf("num2 = %v\n", num2)

	num1 = 24
	fmt.Printf("num1 = %v\n", num1)
	fmt.Printf("num2 = %v\n", num2)
	fmt.Printf("num2 = %v\n", *num2)

}
