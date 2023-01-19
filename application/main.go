package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const cardNumber = "123456789000"
const pinNumber = "123456"

var bankBalance int64 = 100000

var reader = bufio.NewReader(os.Stdin)

func withdraw() {

	fmt.Println("Please enter the amount in multiples of 100! ")
	amount, _ := reader.ReadString('\n')

	withdrawAmount, error := strconv.ParseInt(strings.TrimSpace(amount), 10, 64)
	if error != nil {
		fmt.Println("Invalid amount entered")
		return
	}

	if withdrawAmount > 20000 {
		fmt.Println("You can withdraw maximum 20000 at a time")
		return
	}

	if (withdrawAmount%100 == 0 || withdrawAmount%500 == 0) && withdrawAmount <= bankBalance && withdrawAmount > 0 {
		bankBalance -= withdrawAmount
		fmt.Println("Please collect your cash", withdrawAmount)

		fmt.Println("Do you want to check your remaining balance? (Y/N)")
		checkBalanceChoiceInput, _ := reader.ReadString('\n')
		checkBalanceChoice := strings.TrimSpace(checkBalanceChoiceInput)

		if checkBalanceChoice == "Y" || checkBalanceChoice == "y" {
			fmt.Println("Your remaining balance is", bankBalance)
		}

	} else if withdrawAmount > bankBalance {
		defer tryToRecover()
		fmt.Println("You don't have enough balance")
	} else {
		defer tryToRecover()
		panic("Invalid amount entered!")
	}
}

func cardCheck() bool {

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading card number")
		return false
	}
	cardInput := strings.TrimSpace(input)

	if len(cardInput) != 12 && cardInput != cardNumber {
		defer tryToRecover()
		panic("Invalid card number")
	} else {
		fmt.Println("Please enter your pin")

		input, _ := reader.ReadString('\n')
		pinInput := strings.TrimSpace(input)

		if len(pinInput) == 6 && pinInput == pinNumber {
			return true
		} else {
			fmt.Println("Invalid pin")
			return false
		}
	}
}

func tryToRecover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func deposit() {
	fmt.Println("Please enter the amount you want to deposit: ")
	amountInt, errs := reader.ReadString('\n')
	if errs != nil {
		fmt.Println("Error reading amount")
	}

	amount, amountError := strconv.ParseInt(strings.TrimSpace(amountInt), 10, 64)
	if amountError != nil {
		defer tryToRecover()
		panic("Invalid amount entered")
	}
	if amount < 0 || amount > 100000 {
		fmt.Println("Invalid amount entered")
		return
	}

	bankBalance += amount
	fmt.Println("Your updated balance is", bankBalance)
}

func checkBalance() {
	fmt.Println("Your balance is", bankBalance)
}

func exitProgram() {
	fmt.Println()
	fmt.Println("Thank you for using Simform Bank ATM")
	os.Exit(0)
}

func main() {
	defer fmt.Println("Thank you for using Simform Bank ATM")

	fmt.Println("Welcome to Simform Bank ATM")
	fmt.Println("Please insert your card")

	if !cardCheck() {
		fmt.Println("Thank you for using Simform Bank ATM")
		os.Exit(0)
	}

start:
	fmt.Println()
	fmt.Println("Select one of the following options")
	fmt.Println("1. Withdraw")
	fmt.Println("2. Deposit")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit")
	fmt.Println("Please enter your choice: ")

	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading choice")
	}
	choiceInput := strings.TrimSpace(choice)

	switch choiceInput {
	case "1":
		withdraw()
	case "2":
		deposit()
	case "3":
		checkBalance()
	case "4":
		exitProgram()
	default:
		fmt.Println("Invalid choice")
	}

	fmt.Println("Do you wish to continue? (Y/N)")
	continueChoice, errs := reader.ReadString('\n')

	if errs != nil {
		fmt.Println("Error reading choice")
	}
	continueChoiceInput := strings.TrimSpace(continueChoice)

	if continueChoiceInput == "Y" || continueChoiceInput == "y" {
		time.Sleep(2 * time.Second)
		goto start
	} else {
		fmt.Println()
		fmt.Println("Thank you for using Simform Bank ATM")
		os.Exit(0)
	}

}
