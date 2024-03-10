package main

import (
	"fmt"
)

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func main() {

	var side1, side2, side3 int
	var result bool

	fmt.Scan(&side1, &side2, &side3)

	result = isTriangle(side1, side2, side3)

	if result {
		fmt.Println("segitiga")
	} else {
		fmt.Println("bukan segitiga")
	}
}
