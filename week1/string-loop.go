package main

import "fmt"

func stringLoop(x int, y string) {
	var i int

	for i = 1; i <= x; i++ {
		fmt.Println(y)
	}
}

func main() {
	var x int
	var y string

	fmt.Scan(&x)
	fmt.Scan(&y)

	stringLoop(x, y)
}
