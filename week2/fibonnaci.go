package main

import "fmt"

func main() {
	var num int
	var total int = 0 + 1

	fmt.Scan(&num)

	fibonnaci(num, &total)
	fmt.Println(total)
}

func fibonnaci(x int, total *int) {
	var i, beforeTotal, afterTotal int

	if x != 0 && x == 1 {
		*total = 0
	} else {
		beforeTotal = 1
		afterTotal = 0
		for i = 2; i != x; i++ {
			afterTotal = beforeTotal + *total
			beforeTotal = *total
			*total = afterTotal
		}
	}
}
