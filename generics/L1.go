package main

import (
	"fmt"
)

func getLast[T any](s []T) T {
	var val T;
	if len(s)<=0{
		return val;
	}
	return s[len(s)-1];
}

func main() {
	// Test with a slice of ints
	ints := []int{1, 2, 3, 4}
	lastInt := getLast(ints)
	fmt.Println("Last int:", lastInt) // Output: 4

	// Test with a slice of strings
	strs := []string{"mail", "campaign", "blast"}
	lastStr := getLast(strs)
	fmt.Println("Last string:", lastStr) // Output: blast

	// Test with an empty slice
	empty := []float64{}
	lastFloat := getLast(empty)
	fmt.Println("Last float from empty slice:", lastFloat) // Output: 0
}