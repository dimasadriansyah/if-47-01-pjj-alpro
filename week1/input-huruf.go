package main

import "fmt"

func main() {
	var input string

	found := false

	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		letter := string(input[i])
		if letter == "k" || letter == "q" {
			found = true
			break
		}
	}

	fmt.Print(found)
}
