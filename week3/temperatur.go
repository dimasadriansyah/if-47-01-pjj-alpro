package main

import "fmt"

func main() {
	var c float64

	fmt.Scan(&c)
	temperatur(&c)
}

func temperatur(x *float64) {
	var r, f, k float64

	r = *x * 0.8
	f = (*x * 9 / 5) + 32
	k = *x + 273.15

	fmt.Println(r, "R", f, "F", k, "K")
}
