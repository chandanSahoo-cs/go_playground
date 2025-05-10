package main

import "fmt"

func printPrimes (max int){
	for i:=2;i<=max;i++{
		flag:=true
		for j:=2;j*j<i;j++{
			if i%j==0{
				flag=false;
				break;
			}
		}
		if(flag){
			fmt.Println(i)
		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
