package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array class")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Grape"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"Carrot", "Potato", "Tomato"}

	fmt.Println(vegList)
	fmt.Println(len(vegList))

	// Multi Dimensional ARRAYS
	var twoDArray [4][6]int
	var multiDArray = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	twoDArray[0][0] = 1
	fmt.Println(twoDArray)

	multiDArray[1][2] = 9
	fmt.Println(multiDArray)

}
