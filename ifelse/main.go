package main

import "fmt"

func main() {

	fmt.Println("Welcome to the ifelse class")

	Number := 15

	if Number > 15 {
		fmt.Println("Number is greater than 15")
	} else if Number < 15 {
		fmt.Println("Number is less than 15")
	} else {
		fmt.Println("Number is same")
	}

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than 10")
	}

}
