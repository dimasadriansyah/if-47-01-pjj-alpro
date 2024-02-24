package main

import "fmt"

func tangki(t int) bool {
	var isiTangki, v int
	var result bool

	isiTangki = 0

	for isiTangki <= t {
		fmt.Scan(&v)
		isiTangki = isiTangki + v

		result = isiTangki >= t
		fmt.Println(result)
	}

	return result
}

func main() {
	var t int

	fmt.Scan(&t)
	tangki(t)
}
