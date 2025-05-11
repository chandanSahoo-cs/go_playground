/*
Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)

count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.
package main
*/
package main
import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	mp := make(map[string]bool)
	for _,val :=range(messages){
		words := strings.Fields(val);
		for _,w := range(words){
			mp[strings.ToLower(w)] = true;
		}
	}

	return len(mp);
}

func main() {
	messages := []string{"Hello world", "hello there", "General Kenobi"}
	count := countDistinctWords(messages)
	fmt.Println("Distinct word count:", count) // Output: 5
}


/*
Sample Output :
Distinct word count: 5
*/

