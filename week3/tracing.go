package main

import "fmt"

func main() {
	var a int
	a = 10
	procB(a, a)
	fmt.Print(a)
}

func procB(b, c int) {
	b = b + c
	c = a + b + c
}
