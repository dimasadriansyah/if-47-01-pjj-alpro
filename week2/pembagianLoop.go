package main

import "fmt"

func main() {

	var n, m int

	fmt.Scan(&n, &m)
	fmt.Println(perkalianLoop(n, m))

}

func pembagianLoop(n, m int) int {
	var total int

	total = 0

	if n >= 0 && m > 0 {
		for n >= m {
			n = n - m
			total = total + 1
		}
	}

	return total
}
