/*
Implement the Formatter interface with a method Format that returns a formatted string.
Define structs that satisfy the Formatter interface: PlainText, Bold, Code.
The structs must all have a message field of type string
PlainText should return the message as is.
Bold should wrap the message in two asterisks (**) to simulate bold text (e.g., message).
Code should wrap the message in a single backtick (`) to simulate code block (e.g., message)
*/

package main

import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct{
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (pt PlainText) Format() string {
	return pt.message
}

func (b Bold) Format() string {
	return fmt.Sprintf("**%s**",b.message)
}

func (c Code) Format() string {
	return fmt.Sprintf("`%s`",c.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

func main() {
	formats := []Formatter{
		PlainText{message: "Hello, world!"},
		Bold{message: "Hello, world!"},
		Code{message: "Hello, world!"},
	}

	for _, f := range formats {
		fmt.Println(SendMessage(f))
	}
}

/*
Sample Output:
Hello, world!
**Hello, world!**
`Hello, world!`
*/