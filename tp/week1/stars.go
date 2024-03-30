package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan asli positif ganjil: ")
	fmt.Scan(&n)

	// Memastikan n adalah bilangan bulat positif ganjil
	if n <= 0 || n%2 == 0 {
		fmt.Println("Masukkan harus merupakan bilangan bulat positif ganjil.")
		return
	}

	// Mencetak pola bintang
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i || j == n-1-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
