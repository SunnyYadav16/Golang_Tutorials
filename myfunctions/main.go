package main

import "fmt"

func main() {

	fmt.Println("Welcome to the functions class")
	greeter()

	value := adder(10, 20)
	fmt.Println("Value is: ", value)

	proValue := proAdder(10, 2, 20, 34, 213, 34)
	fmt.Println("Pro value is: ", proValue)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("Hello World from greeter")
}
