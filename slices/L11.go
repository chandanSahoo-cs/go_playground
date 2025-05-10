/*
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice.
Ensure the tags field contains an initialized slice. No nil slices.
Complete the tagger function. It should take an sms message and return a slice of strings.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
*/
package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i,msg := range(messages) {
		messages[i].tags = tagger(msg)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	
	words := strings.Fields(msg.content)

	for _,w := range(words) {
		if strings.ToLower(w)=="urgent"{
			tags = append(tags, "Urgent")
		}
		if strings.ToLower(w)=="sale"{
			tags = append(tags, "Promo")
		}
	}
	return tags
}

func main() {
	messages := []sms{
		{id: "1", content: "Huge URGENT SALE now"},
		{id: "2", content: "Meeting reminder"},
		{id: "3", content: "Limited time sale offer"},
	}

	tagged := tagMessages(messages, tagger)

	for _, m := range tagged {
		fmt.Printf("ID: %s, Tags: %v\n", m.id, m.tags)
	}
}

/*
Sample Output:
ID: 1, Tags: [Urgent Promo]
ID: 2, Tags: []
ID: 3, Tags: [Promo]
*/
