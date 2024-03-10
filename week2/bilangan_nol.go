package main

import "fmt"

func main() {
	var bilangan int
	var jenis bool

	fmt.Scan(&bilangan)
	jenis = nol(bilangan)
	fmt.Println(jenis)
}

func nol(x int) bool {
	var isNol bool

	if x == 0 {
		isNol = true
	} else {
		isNol = false
	}

	return isNol
}
