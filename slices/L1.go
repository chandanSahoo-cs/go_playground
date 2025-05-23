/*
When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

Complete the getMessageWithRetries function. It takes three strings and returns:

An array of 3 strings
An array of 3 integers
The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

"hello" costs 5
"world" costs 5, adding "hello" makes total cost 10 (5 + 5)
"!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)
*/

package main
import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages:=[3]string{primary,secondary,tertiary}
	costs:=[3]int{len(primary),len(primary+secondary),len(primary+secondary+tertiary)}

	return messages,costs
}

func main() {
	msgs, costs := getMessageWithRetries("hello", "world", "!")
	fmt.Println("Messages:", msgs)
	fmt.Println("Costs:", costs)
}

/*
Sample Output
Messages: [hello world !]
Costs: [5 10 11]
*/

