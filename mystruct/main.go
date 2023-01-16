package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Struct tutorial")

	my_user_struct := User{"John", "happy@gomail.com", true, 20}
	fmt.Println("Struct details are: ", my_user_struct)
	fmt.Printf("Struct details are: %+v", my_user_struct)
	fmt.Printf("Name is %v and email is %v", my_user_struct.Name, my_user_struct.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
