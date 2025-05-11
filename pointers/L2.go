/*
Write a getMessageText function. It should accept a pointer to an Analytics struct and a Message struct (not a pointer).

It should look at the Success field of the Message struct and, based on that, increment the MessagesTotal, MessagesSucceeded, or MessagesFailed fields of the Analytics struct as appropriate.
*/


package main

import (
	"fmt"
)

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(a *Analytics,m Message){
	if(m.Success){
		a.MessagesSucceeded++
	}else{
		a.MessagesFailed++
	}
	a.MessagesTotal++
}

func main() {
	analytics := &Analytics{}
	messages := []Message{
		{"user1", true},
		{"user2", false},
		{"user3", true},
		{"user4", false},
		{"user5", true},
	}

	for _, m := range messages {
		getMessageText(analytics, m)
	}

	fmt.Println("Analytics Summary:")
	fmt.Println("Total:", analytics.MessagesTotal)
	fmt.Println("Succeeded:", analytics.MessagesSucceeded)
	fmt.Println("Failed:", analytics.MessagesFailed)
}



/*
Sample Output :
Analytics Summary:
Total: 5
Succeeded: 3
Failed: 2
*/
