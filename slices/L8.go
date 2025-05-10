/*
We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.

For example, a 5x10 matrix, produced from calling createMatrix(5, 10), would look like this:

[0  0  0  0  0  0  0  0  0  0]
[0  1  2  3  4  5  6  7  8  9]
[0  2  4  6  8 10 12 14 16 18]
[0  3  6  9 12 15 18 21 24 27]
[0  4  8 12 16 20 24 28 32 36]
*/


package main

import (
	"fmt"
)

func createMatrix(rows, cols int) [][]int {
	ans := [][]int{}

	for i:=0;i<rows;i++{
		r := make([]int,cols)
		for j:=0;j<cols;j++{
			r[j]=i*j;
		}
		ans = append(ans, r)
	}

	return ans
}

func main() {
	// Create a 5x10 matrix
	matrix := createMatrix(5, 10)

	// Print the matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
}

/* 
Sample Output:
[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 4 5 6 7 8 9]
[0 2 4 6 8 10 12 14 16 18]
[0 3 6 9 12 15 18 21 24 27]
[0 4 8 12 16 20 24 28 32 36]
*/