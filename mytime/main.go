package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time tutorial")
	presentTime := time.Now()

	fmt.Println("The time is", presentTime)

	formatted_time := presentTime.Format("01-02-2006")
	fmt.Println("Formatted Time", formatted_time)

	createdDate := time.Date(2020, time.August, 16, 20, 20, 0, 0, time.UTC)
	fmt.Println("Created Date", createdDate)

}
