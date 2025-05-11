/*
We can speed up our contact-info lookups by using a map!

Key-based map lookup: O(1)
Slice brute-force search: O(n)
Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user's name and phone number. The first name in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/

package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	ansMap :=make(map[string]user)
	if(len(names)!=len(phoneNumbers)){
		return nil, errors.New("invalid sizes")
	}

	for i:=0;i<len(names);i++{
		ansMap[names[i]] = user{
			name : names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return ansMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	numbers := []int{1234, 5678, 9012}

	usersMap, err := getUserMap(names, numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for name, userData := range usersMap {
		fmt.Printf("Name: %s, Phone: %d\n", name, userData.phoneNumber)
	}
}

/*
Sample Output : 
Name: Alice, Phone: 1234
Name: Bob, Phone: 5678
Name: Charlie, Phone: 9012
*/
