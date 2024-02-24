package main

import "fmt"

func main() {
	var b1, b2, b3, gr, kg int

	kg = 0

	fmt.Scan(&b1, &b2, &b3)

	gr = b1 + b2 + b3

	for gr >= 1000 {
		gr = gr - 1000
		kg = kg + 1
	}

	fmt.Println(kg, gr)

}
