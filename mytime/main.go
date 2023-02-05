package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time tutorial")
	presentTime := time.Now()

	fmt.Println("The time is", presentTime)

	/*
		We need to provide the same time format layout for changing or formatting the time and date
		formatted_time := presentTime.Format("01-02-2006 15:04:05 Monday")

		To get the current date in the formatted style, we need to use this date: 01-02-2006
		To get the current time in the formatted style, we need to use this time: 15:04:05
		To get the current day in the formatted style, we need to use this day: Monday
	*/
	formattedTime := presentTime.Format("01-02-2006")

	fmt.Println("Formatted Time", formattedTime)

	createdDate := time.Date(2020, time.August, 16, 20, 20, 0, 0, time.UTC)
	fmt.Println("Created Date", createdDate)

}
