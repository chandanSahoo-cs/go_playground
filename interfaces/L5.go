/*
After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the getExpenseReport function using a type switch.

If the expense is an email, return the email's toAddress and the cost of the email.
If the expense is an sms, return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, return an empty string and 0.0 for the cost.
*/

package main
import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch exp := e.(type){
	case email:
		return exp.toAddress,exp.cost()
	case sms:
		return exp.toPhoneNumber,exp.cost()
	default:
		return "",0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	// Define different expenses
	e1 := email{
		isSubscribed: true,
		body:         "Hello, subscriber!",
		toAddress:    "user@example.com",
	}

	e2 := sms{
		isSubscribed: false,
		body:         "Verify your account",
		toPhoneNumber: "+1987654321",
	}

	e3 := invalid{}

	// Slice of expenses
	expenses := []expense{e1, e2, e3}

	// Print report for each expense
	for _, e := range expenses {
		recipient, cost := getExpenseReport(e)
		fmt.Printf("To: %s | Cost: $%.2f\n", recipient, cost)
	}
}

/*
Sample Output
To: user@example.com | Cost: $0.18
To: +1987654321 | Cost: $2.00
To:  | Cost: $0.00
*/
