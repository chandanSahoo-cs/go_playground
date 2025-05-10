/*
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.
*/


package main

import (
	"fmt"
)

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	ans := []Message{}

	for i:=0;i<len(messages);i++ {
		if(messages[i].Type()==filterType){
			ans = append(ans,messages[i])
		}
	}

	return ans;
}

func main() {
	messages := []Message{
		TextMessage{Sender: "Alice", Content: "Hello"},
		MediaMessage{Sender: "Bob", MediaType: "image", Content: "Photo.jpg"},
		LinkMessage{Sender: "Carol", URL: "https://example.com", Content: "Check this out"},
		TextMessage{Sender: "Dave", Content: "What's up?"},
	}

	filtered := filterMessages(messages, "text")

	fmt.Println("Filtered messages of type 'text':")
	for _, m := range filtered {
		fmt.Printf("%#v\n", m)
	}
}

/*
Sample Output:
Filtered messages of type 'text':
main.TextMessage{Sender:"Alice", Content:"Hello"}
main.TextMessage{Sender:"Dave", Content:"What's up?"}
*/
