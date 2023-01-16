package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Struct tutorial")

	my_user_struct := User{"John", "happy@gomail.com", true, 20}
	fmt.Println("Struct details are: ", my_user_struct)
	fmt.Printf("Struct details are: %+v\n", my_user_struct)
	fmt.Printf("Name is %v and email is %v\n", my_user_struct.Name, my_user_struct.Email)
	my_user_struct.GetStatus()
	my_user_struct.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status is: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "sunny@gmail.com"
	fmt.Println("New email is: ", u.Email)
}
