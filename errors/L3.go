/*
Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

can not divide DIVIDEND by zero

Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.
*/

package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string{
	return fmt.Sprintf("can not divide %v by zero",de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	// Test with valid division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Test with divide by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

/*
Sample Output:
Result: 5
Error: can not divide 10 by zero
*/