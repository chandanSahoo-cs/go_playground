/*
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.
*/

package main

import (
	"fmt"
)


func indexOfFirstBadWord(msg []string, badWords []string) int {

	for idx, words := range(msg) {
		for _, badWrds :=range(badWords){
			if words==badWrds{
				return idx
			}
		}
	}
	return -1
}

func main() {
	msg := []string{"hello", "this", "is", "bad", "stuff"}
	badWords := []string{"bad", "ugly", "nasty"}

	index := indexOfFirstBadWord(msg, badWords)
	fmt.Println("First bad word index:", index) // Output: 3

	index = indexOfFirstBadWord(msg, []string{"mean"})
	fmt.Println("First bad word index:", index) // Output: -1
}

/*
Sample Output: 
First bad word index: 3
First bad word index: -1
*/
