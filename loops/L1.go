/*
At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the bulkSend() function.

It should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
...
Use a loop to calculate the total cost and return it.
*/
package main

import "fmt"

func bulkSend(numMessages int) float64{
	totalCost:=0.0
	for i:=0;i<numMessages;i++{
		totalCost+=1+float64(i)*0.01
	}
	return totalCost
}

func main() {
	// Test 1: Sending 0 messages
	fmt.Println("Cost for 0 messages:", bulkSend(0)) // Expected: 0.0

	// Test 2: Sending 1 message
	fmt.Println("Cost for 1 message:", bulkSend(1)) // Expected: 1.0

	// Test 3: Sending 5 messages
	fmt.Println("Cost for 5 messages:", bulkSend(5))
	// Breakdown:
	// i = 0 → 1.00
	// i = 1 → 1.01
	// i = 2 → 1.02
	// i = 3 → 1.03
	// i = 4 → 1.04
	// Total: 5.10

	// Test 4: Sending 10 messages
	fmt.Println("Cost for 10 messages:", bulkSend(10)) // Expected: 10.45
}

/*
Sample Output:
Cost for 0 messages: 0
Cost for 1 message: 1
Cost for 5 messages: 5.1
Cost for 10 messages: 10.45
*/