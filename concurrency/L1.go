package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
}

// Don't touch below this line

func test(message string) {
	go sendEmail(message)
	fmt.Printf("Email sent: '%s'\n", message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

// func main() {
// 	test("Hello there Kaladin!")
// 	test("Hi there Shallan!")
// 	test("Hey there Dalinar!")
// }
