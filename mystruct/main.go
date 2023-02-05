package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Struct tutorial")

	myUserStruct := User{"John", "happy@gomail.com", true, 20}
	fmt.Println("Struct details are: ", myUserStruct)
	fmt.Printf("Struct details are: %+v", myUserStruct)
	fmt.Printf("Name is %v and email is %v", myUserStruct.Name, myUserStruct.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
