package main

import "fmt"

func main() {
	var alfaLow, alfaUp byte
	fmt.Scanf("%c", &alfaUp)
	alfaLow = upperToLow(alfaUp)
	fmt.Printf("%c", alfaLow)
}

func upperToLow(char byte) byte {
	return char + 32
}
