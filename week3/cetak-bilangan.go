package main

import "fmt"

func main() {
	var x, y, a, b int
	fmt.Scan(&x, &y, &a, &b)
	cetakBilangan(x, y, a, b)
}

func cetakBilangan(x, y, a, b int) {

	var j, k int

	for x <= y {

		j = (x / 100) % 10
		k = (x / 1) % 10

		if j != a && k != b {
			fmt.Println(x)
		}
		x = x + 1
	}
}
