package main

import "fmt"

func menghitung(berat int) {
	var kg, gr, hargaKg, hargaGr int
	kg = 0
	gr = 0

	for berat >= 1000 {
		berat = berat - 1000
		kg = kg + 1
	}

	hargaKg = kg * 10000

	gr = gr + berat
	if gr >= 500 && kg < 10 {
		hargaGr = gr * 5
	} else if gr < 500 && kg < 10 {
		hargaGr = gr * 15
	}

	fmt.Println(kg, "kg", "+", gr, "gr")
	fmt.Println("Rp.", hargaKg, "+", "Rp.", hargaGr, "= Rp.", hargaKg+hargaGr)
}

func main() {
	var berat int

	fmt.Scan(&berat)
	menghitung(berat)
}
