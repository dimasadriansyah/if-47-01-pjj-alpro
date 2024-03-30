package main

import "fmt"

func main() {

	var x, y float64
	fmt.Scan(&x, &y)

	var kuadran string
	if x > 0 && y > 0 {
		kuadran = "Kuadran I"
	} else if x < 0 && y > 0 {
		kuadran = "Kuadran II"
	} else if x < 0 && y < 0 {
		kuadran = "Kuadran III"
	} else if x > 0 && y < 0 {
		kuadran = "Kuadran IV"
	} else {
		kuadran = "Pada sumbu atau pusat"
	}
	fmt.Println(kuadran)
}
