package main

import (
	"fmt"
)

func fibonacciSum(n int) int {
	var sum, i int
	var fib []int

	if n <= 0 {
		return 0
	}

	fib = []int{0, 1}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	sum = 0
	for i = 0; i < n; i++ {
		sum += fib[i]
	}

	return sum
}

func main() {
	var n, result int

	fmt.Scan(&n)

	result = fibonacciSum(n)

	fmt.Println(result)
}
