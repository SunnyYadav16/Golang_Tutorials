package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers class")

	var ptr *int
	fmt.Println("Value of ptr is", ptr)

	myNumber := 42

	ptr = &myNumber
	fmt.Println("Value address of ptr is", ptr)
	fmt.Println("Value of ptr is", *ptr)

	*ptr += 2
	fmt.Println("Value of ptr is", myNumber)
}
