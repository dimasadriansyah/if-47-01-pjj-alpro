package main

import "fmt"

func f(x, y int) float64 {
	var hasilKali, i int

	hasilKali = 1

	for i = 1; i <= 2; i++ {
		hasilKali *= x
	}
	return (1.0)/(3.0*float64(hasilKali)+10.0) + (10.0 * float64(y)) + (7.0)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(f(x, y))
}
