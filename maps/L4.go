/*
Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe

Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
*/


package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	ansMap := make(map[rune]map[string]int)

	for _,val := range(names){
		if _,ok :=ansMap[rune(val[0])]; !ok{
			mn := make(map[string]int)
			ansMap[rune(val[0])] = mn
		}

		if _,ok :=ansMap[rune(val[0])][val] ; !ok{
			ansMap[rune(val[0])][val] = 0
		}

		ansMap[rune(val[0])][val]++;
	}

	return ansMap;
}

func main() {
	names := []string{"billy", "billy", "bob", "joe", "john", "jane", "bob", "jack"}
	result := getNameCounts(names)

	for firstChar, nameMap := range result {
		fmt.Printf("%c:\n", firstChar)
		for name, count := range nameMap {
			fmt.Printf("  %s: %d\n", name, count)
		}
	}
}

/*
Sample Output : 
b:
  billy: 2
  bob: 2
j:
  joe: 1
  john: 1
  jane: 1
  jack: 1
*/