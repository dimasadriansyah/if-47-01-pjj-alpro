package main

import "fmt"

func main() {
	// Jumlah uang dalam koin Lumin
	var lumin, uangLumin, quantumShard, sisaUang, galactum, nebulaCrown, stellaris int

	fmt.Scan(&uangLumin)
	// Menghitung koin Quantum Shard
	quantumShard = uangLumin / (10 * 3 * 2 * 20)
	sisaUang = uangLumin % (10 * 3 * 2 * 20)

	// Menghitung koin Galactum dari sisa uang
	galactum = sisaUang / (3 * 2 * 20)
	sisaUang = sisaUang % (3 * 2 * 20)

	// Menghitung koin Nebula Crown dari sisa uang
	nebulaCrown = sisaUang / (2 * 20)
	sisaUang = sisaUang % (2 * 20)

	// Menghitung koin Stellaris dari sisa uang
	stellaris = sisaUang / 20

	// Sisa uang setelah dikonversi ke Stellaris adalah uang Lumin
	lumin = sisaUang % 20

	fmt.Println(quantumShard, galactum, nebulaCrown, stellaris, lumin)
}
