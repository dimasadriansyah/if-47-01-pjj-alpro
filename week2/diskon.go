package main

import "fmt"

func main() {
	var x, price float64
	var isMember bool

	fmt.Scan(&x, &isMember)
	price = diskon(x, isMember)
	fmt.Println(price)
}

func diskon(biaya float64, isMember bool) float64 {
	var discount float64

	discount = 0

	if biaya > 100000 && isMember {
		discount = biaya * float64(0.1)
		biaya = biaya - discount
	} else if biaya > 100000 && !isMember {
		discount = biaya * float64(0.05)
		biaya = biaya - discount
	}

	return biaya
}
