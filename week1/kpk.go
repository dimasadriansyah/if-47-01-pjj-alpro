package main

import (
	"fmt"
)

func findKPK(n, m int) int {
	var kpk int

	kpk = n

	for kpk%m != 0 {
		kpk += n
	}

	return kpk
}

func main() {

	var n, m int

	fmt.Scan(&n, &m)

	result := findKPK(n, m)
	fmt.Printf("KPK dari %d dan %d adalah: %d\n", n, m, result)
}
