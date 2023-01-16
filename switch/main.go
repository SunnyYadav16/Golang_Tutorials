package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch class")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(7)
	fmt.Println("Dice Number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1")
	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("You got 3")
		fallthrough
	case 4:
		fmt.Println("You got 4")
		fallthrough
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	}

}
