/*
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent

COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient's phone number
Be sure to include the $ symbol and the single quotes

*/


package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent",cost,recipient)
}

func main() {
	fmt.Println(getSMSErrorString(2.5, "+1234567890"))
	fmt.Println(getSMSErrorString(10, "9876543210"))
}

/*
SMS that costs $2.50 to be sent to '+1234567890' can not be sent
SMS that costs $10.00 to be sent to '9876543210' can not be sent
*/

