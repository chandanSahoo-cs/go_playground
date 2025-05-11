/*
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

Word	Replacement
fubb	****
shiz	****
witch	*****
It should update the value in the pointer and return nothing.

Do not alter the function signature.
*/

package main

import (
	"strings"
	"fmt"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message,"fubb","****")
	*message = strings.ReplaceAll(*message,"shiz","****")
	*message = strings.ReplaceAll(*message,"witch","*****")
	return
}

func main() {
	messages := []string{
		"This is a fubb message",
		"Don't say shiz!",
		"She's a witch",
		"Clean message here",
		"Multiple fubb and shiz in fubbshizwitch",
	}

	for _, msg := range messages {
		fmt.Print("Original: ", msg, "\nUpdated:  ")
		removeProfanity(&msg)
		fmt.Println(msg)
		fmt.Println()
	}
}

/*
Sample Output:
Original: This is a fubb message
Updated:  This is a **** message

Original: Don't say shiz!
Updated:  Don't say ****!

Original: She's a witch
Updated:  She's a *****

Original: Clean message here
Updated:  Clean message here

Original: Multiple fubb and shiz in fubbshizwitch
Updated:  Multiple **** and **** in **************
*/