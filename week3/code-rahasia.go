package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)
	hitungJumlah(n, &jumlah)
	fmt.Println(jumlah)
}

func hitungJumlah(n int, jumlah *int) {
	var i, j, k, l int

	j = 0
	k = 0
	l = n

	for i = 0; i < l; i++ {
		fmt.Scan(&n)
		j = (n / 1000) % 10
		k = (n / 10) % 10
		*jumlah = *jumlah + j + k
	}
}
