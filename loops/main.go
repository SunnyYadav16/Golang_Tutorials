package main

import "fmt"

func main() {

	fmt.Println("Welcome to the loops class")

	days := []string{"Sunday", "Monday", "Wednesday", "Saturday"}

	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	for index, day := range days {
		fmt.Println("Index:", index, "Day:", day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto statement
		}
		//if rogueValue == 5 {
		//	rogueValue++
		//	continue
		//}

		//if rogueValue == 5{
		//	break
		//}

		fmt.Println("Value: ", rogueValue)
		rogueValue++
	}

statement:
	fmt.Println("In the goto statement")
}
