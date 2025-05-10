/*
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.
*/

package main

import (
	"fmt"
)

func sum(nums ...int) int {
	ans:=0
	for i:=0;i<len(nums);i++ {
		ans+=nums[i];
	}

	return ans;
}

func main() {
	// Direct variadic call
	fmt.Println(sum(1, 2, 3)) // Output: 6

	// Using a slice with spread operator
	nums := []int{4, 5, 6}
	fmt.Println(sum(nums...)) // Output: 15
}


/*
Sample Output:
Sum: 6
Sum: 15
*/