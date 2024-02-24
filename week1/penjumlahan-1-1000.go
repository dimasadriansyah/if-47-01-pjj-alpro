package main

import "fmt"

func main() {

	sum := 0

	i := 1
	for i <= 1000 {

		sum += i

		i++
	}

	fmt.Print(sum)
}
