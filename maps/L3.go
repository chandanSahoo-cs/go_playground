/*
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the getCounts function. It takes as input:

messagedUsers: a slice of strings.
validUsers: a map of string -> int.
It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/


package main

import (
	"fmt"
)

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _,val := range(messagedUsers){
		if _,ok :=validUsers[val] ; ok {
			validUsers[val]++
		}
	}
}


func main() {
	messagedUsers := []string{"alice", "bob", "alice", "dave", "bob", "carol", "bob"}
	validUsers := map[string]int{
		"alice": 0,
		"bob":   0,
		"carol": 0,
	}

	getCounts(messagedUsers, validUsers)

	fmt.Println(validUsers)
}

/*
Sample Output: 
map[alice:2 bob:3 carol:1]
*/

