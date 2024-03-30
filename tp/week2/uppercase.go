package main

import "fmt"

func lowT0Upper(kar byte) byte {
	if kar >= 'a' && kar <= 'z' {
		return kar - 32
	}

	return kar
}

func main() {
	var lowercaseChar, uppercaseChar byte

	fmt.Scanf("%c", &lowercaseChar)

	uppercaseChar = lowT0Upper(lowercaseChar)
	fmt.Println(string(uppercaseChar))
}
