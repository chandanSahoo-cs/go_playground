/*Complete the validateStatus function. It should return an error when the status update violates any of the rules:

If the status is empty, return an error that reads status cannot be empty.
If the status exceeds 140 characters, return an error that says status exceeds 140 characters.
*/

package main

import "fmt"

func validateStatus(status string) error {
	if len(status)>140 {
		return fmt.Errorf("status exceeds 140 characters")
	}else if len(status)==0{
		return fmt.Errorf("status cannot be empty")
	}

	return nil
}

func main() {
	testCases := []string{
		"", // too short
		"Hello, world!", // valid
		string(make([]byte, 141)), // too long
	}

	for i, status := range testCases {
		err := validateStatus(status)
		if err != nil {
			fmt.Printf("Test %d: Error - %s\n", i+1, err)
		} else {
			fmt.Printf("Test %d: Status is valid\n", i+1)
		}
	}
}
/*
Test 1: Error - status cannot be empty
Test 2: Status is valid
Test 3: Error - status exceeds 140 characters
*/


