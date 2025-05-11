/*
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
*/


package main

import (
	"fmt"
	"errors"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	userEntry,ok := users[name];

	if !ok {
		return false,errors.New("not found")
	}

	if !userEntry.scheduledForDeletion{
		return false,nil
	}

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	users := map[string]user{
		"Alice":   {name: "Alice", number: 1234, scheduledForDeletion: false},
		"Bob":     {name: "Bob", number: 5678, scheduledForDeletion: true},
		"Charlie": {name: "Charlie", number: 9012, scheduledForDeletion: false},
	}

	deleted, err := deleteIfNecessary(users, "Bob")
	fmt.Println("Deleted:", deleted, "Error:", err) // Expected: Deleted: true Error: <nil>

	deleted, err = deleteIfNecessary(users, "Alice")
	fmt.Println("Deleted:", deleted, "Error:", err) // Expected: Deleted: false Error: <nil>

	deleted, err = deleteIfNecessary(users, "David")
	fmt.Println("Deleted:", deleted, "Error:", err) // Expected: Deleted: false Error: not found
}

/*
Sample Output : 
Deleted: true Error: <nil>
Deleted: false Error: <nil>
Deleted: false Error: not found
*/
