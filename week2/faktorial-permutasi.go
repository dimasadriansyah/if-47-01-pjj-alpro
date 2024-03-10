package main

import "fmt"

func f(x int) int {
	var xHasil, i int

	if x == 0 {
		return 1
	}

	xHasil = 1

	for i = 1; i != x; {
		xHasil *= x
		x = x - 1
	}

	return xHasil
}

func main() {
	var x, y, result1, result2, result3 int

	fmt.Scan(&x, &y)

	result1 = f(x)
	result2 = f(y)

	if x >= y {
		result3 = f(x) / f(x-y)
	} else if x < y {
		result3 = f(y) / f(y-x)
	}

	fmt.Println(result1, result2, result3)

}
