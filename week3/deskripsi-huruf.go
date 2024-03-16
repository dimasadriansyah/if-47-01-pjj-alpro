package main

import "fmt"

func main() {
	var huruf byte
	fmt.Scanf("%c", &huruf)
	deskripsiHuruf(huruf)
}

func deskripsiHuruf(huruf byte) {

	var angka int

	if huruf == 'A' {
		angka = 5
	} else if huruf == 'B' {
		angka = 4
	} else if huruf == 'C' {
		angka = 3
	} else if huruf == 'D' {
		angka = 2
	} else if huruf == 'E' {
		angka = 1
	} else if huruf == 'T' {
		angka = 0
	}

	fmt.Println(angka)
}
