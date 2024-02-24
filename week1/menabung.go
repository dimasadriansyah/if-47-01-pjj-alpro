package main

import "fmt"

func main() {
	var x, y, i, total int

	total = 0

	fmt.Scan(&x)

	if x == 4 {
		for i = 1; i <= x; i++ {
			fmt.Scan(&y)
			total = total + y
		}
	} else if x == 5 {
		for i = 1; i <= x; i++ {
			fmt.Scan(&y)
			total = total + y
		}
	}

	fmt.Print(total)
}
