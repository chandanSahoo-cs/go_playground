/*
The birthdayMessage and sendingReport structs already have implementations of the getMessage method. The getMessage method returns a string, and any type that implements the method can be considered a message (meaning it implements the message interface).

Add the getMessage() method signature as a requirement on the message interface.
Complete the sendMessage function. It should return:
The content of the message.
The cost of the message, which is the length of the message multiplied by 3.
Notice that your code doesn't care at all about whether a specific message is a birthdayMessage or a sendingReport!
*/

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	return msg.getMessage(),len(msg.getMessage())
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	// Test with birthdayMessage
	bm := birthdayMessage{
		birthdayTime:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
		recipientName: "Chandan",
	}
	message1, length1 := sendMessage(bm)
	fmt.Println("Birthday Message:", message1)
	fmt.Println("Length:", length1)

	fmt.Println()

	// Test with sendingReport
	sr := sendingReport{
		reportName:    "Monthly Analytics",
		numberOfSends: 150,
	}
	message2, length2 := sendMessage(sr)
	fmt.Println("Report Message:", message2)
	fmt.Println("Length:", length2)
}

/*
Output Example: 
Birthday Message: Hi Chandan, it is your birthday on 2025-06-10T00:00:00Z
Length: 60

Report Message: Your "Monthly Analytics" report is ready. You've sent 150 messages.
Length: 73
*/

