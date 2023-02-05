package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices class")

	var fruitList = []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Printf("Fruit List Type: %T\n", fruitList)

	fruitList = append(fruitList, "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	scores := make([]int, 3)
	fmt.Println(scores)

	scores[0] = 20
	//scores[1] = 10
	scores[2] = 30
	//scores[3] = 40               // This will throw an error because we have only defined slice upto 3 places
	fmt.Println(scores)
	fmt.Println("Length: ", len(scores))

	scores = append(scores, 40)
	fmt.Println(len(scores))

	// SORTING the SLICE
	fmt.Println(sort.IntsAreSorted(scores))

	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// Delete values from slice based on index
	var courses = []string{"Docker", "Kubernetes", "Puppet", "Terraform", "AWS", "Azure", "GCP"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
