package main

import (
	"time"
	"fmt"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

// func main() {
// 	// Create test emails with different dates
// 	emails := [3]email{
// 		{
// 			body: "Old email",
// 			date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC), // before 2020
// 		},
// 		{
// 			body: "New email",
// 			date: time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC), // after 2020
// 		},
// 		{
// 			body: "Another old email",
// 			date: time.Date(2015, time.July, 15, 0, 0, 0, 0, time.UTC), // before 2020
// 		},
// 	}

// 	// Call the checkEmailAge function
// 	results := checkEmailAge(emails)

// 	// Print the results
// 	for i, res := range results {
// 		fmt.Printf("Email %d ('%s') is old? %t\n", i+1, emails[i].body, res)
// 	}
// }
