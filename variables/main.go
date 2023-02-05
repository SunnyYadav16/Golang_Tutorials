package main

import "fmt"

const LoginToken string = "1234@asdasdjashdj" // Public variable is declared using First letter Capital

func main() {
	var username string = "gopher"
	//fmt.Print("Hello, " + username)
	//fmt.Printf("Hello, %T", username)
	fmt.Println("Type,", username)

	var isLoggedIn bool = false
	fmt.Println("Hello,", isLoggedIn)
	fmt.Println("Type, %T", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("Hello,", smallVal)
	fmt.Println("Type, %T", smallVal)

	var smallFloat float32 = 255.12332131231232312123
	fmt.Println("Hello,", smallFloat)
	fmt.Println("Type, %T", smallFloat)

	var bigFloat float64 = 255.12332131231232312123
	fmt.Println("Hello,", bigFloat)
	fmt.Println("Type, %T", bigFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Println("Hello,", anotherVariable)
	fmt.Println("Type, %T", anotherVariable)

	// implicit type - Cannot change the type of the variable if declared once
	var website = "abcd.com"
	fmt.Println("Hello,", website)

	// no var style
	numberOfUsers := 100
	fmt.Println("Hello,", numberOfUsers)

	// constants
	fmt.Println("Hello,", LoginToken)
}
