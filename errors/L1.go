/*
We offer a product that allows businesses to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
Do the same for the msgToSpouse
If both messages are sent successfully, return the total cost of the messages added together.
*/

package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mssgCus, err :=sendSMS(msgToCustomer);
	if err!=nil{
		return 0,err
	}

	mssgSp, err := sendSMS(msgToSpouse)

	if err!=nil{
		return 0, err
	}

	return mssgCus+mssgSp,nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}


func main() {
	// Case 1: Successful messages
	fmt.Println("== Test: Success ==")
	totalCost, err := sendSMSToCouple("Love you!", "Dinner at 7?")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total cost:", totalCost)
	}

	// Case 2: Error due to long message
	fmt.Println("\n== Test: Error ==")
	totalCost, err = sendSMSToCouple("This message is way too long and should fail.", "Short one.")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total cost:", totalCost)
	}
}

/*
Sample Output: 
== Test: Success ==
Total cost: 26

== Test: Error ==
Error: can't send texts over 25 characters
*/

