package main

import (
	"fmt"
	"time"
)

func main() {
	// Current date and time
	now := time.Now()
	fmt.Println("Current Date and Time:", now)

	// Formatting a date as a string
	formattedDate := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Date:", formattedDate)

	// Creating a specific date
	specificDate := time.Date(2023, time.March, 15, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific Date:", specificDate)

	// Adding and subtracting time
	oneDay := time.Hour * 24
	oneWeekLater := now.Add(oneDay * 7)
	fmt.Println("One week later:", oneWeekLater)

	// Comparing dates
	if now.Before(specificDate) {
		fmt.Println("Current date is before the specific date.")
	} else if now.After(specificDate) {
		fmt.Println("Current date is after the specific date.")
	} else {
		fmt.Println("Current date is the same as the specific date.")
	}

	// Extracting components of a date
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)
}
