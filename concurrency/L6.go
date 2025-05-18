package main

import (
	"fmt"
)

func concurrentFib(n int) []int {
	
	ch :=make(chan int,n);
	go fibonacci(n,ch)

	ans :=[]int{}
	for v :=range(ch){
		ans = append(ans, v);
	}

	return ans
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	n := 10
	result := concurrentFib(n)
	fmt.Printf("First %d Fibonacci numbers: %v\n", n, result)
}