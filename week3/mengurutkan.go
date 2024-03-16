package main

import "fmt"

func main() {
	var x, y int

	x = 0
	y = 1

	for x != y {
		fmt.Scan(&x, &y)
		mengurutkan(&x, &y)
	}
}

func mengurutkan(x, y *int) {
	if *x < *y {
		fmt.Println(*x, *y)
	} else if *y < *x {
		fmt.Println(*y, *x)
	}
}
