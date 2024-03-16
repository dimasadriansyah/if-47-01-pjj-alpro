package main

import "fmt"

func main() {
	var x, y, z int
	var s string

	fmt.Scan(&x, &y)
	fmt.Scan(&s, &z)

	if s == "+" {
		zoomIn(&x, &y, &z, &s)
	} else if s == "-" {
		zoomOut(&x, &y, &z, &s)
	}
}

func zoomIn(x, y, z *int, s *string) {
	*x = *x * *z
	*y = *y * *z
	fmt.Println(*x, *y)
}

func zoomOut(x, y, z *int, s *string) {
	*x = *x / *z
	*y = *y / *z
	fmt.Println(*x, *y)
}
