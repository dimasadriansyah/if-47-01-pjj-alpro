package main

import "fmt"

func main() {
	var n, kekuatanKenshin, kecepatanKenshin, lawanDikalahkan int
	fmt.Scan(&n)

	fmt.Scan(&kekuatanKenshin, &kecepatanKenshin)

	lawanDikalahkan = 0

	for i := 0; i < n; i++ {
		var kekuatanLawan, kecepatanLawan int
		fmt.Scan(&kekuatanLawan, &kecepatanLawan)

		// Melakukan pertarungan jika kekuatan dan kecepatan lawan serta Kenshin minimal 3
		for kekuatanKenshin >= 3 && kecepatanKenshin >= 3 && kekuatanLawan >= 3 && kecepatanLawan >= 3 {
			// Lawan memiliki kekuatan dan kecepatan yang lebih besar, Kenshin menghindari
			if kekuatanLawan >= kekuatanKenshin && kecepatanLawan >= kecepatanKenshin {
				break
			}
			// Lawan memiliki kekuatan lebih besar tetapi kecepatan lebih kecil, Kenshin menggunakan kecepatan
			if kekuatanLawan >= kekuatanKenshin && kecepatanLawan < kecepatanKenshin {
				kecepatanKenshin -= 6
				lawanDikalahkan++
				break
			}
			// Lawan memiliki kekuatan lebih kecil tetapi kecepatan lebih besar, Kenshin menggunakan kekuatan
			if kekuatanLawan < kekuatanKenshin && kecepatanLawan >= kecepatanKenshin {
				kekuatanKenshin -= 6
				lawanDikalahkan++
				break
			}
			// Lawan memiliki kekuatan dan kecepatan lebih kecil, Kenshin menggunakan sisa yang lebih besar
			if kekuatanKenshin >= kecepatanKenshin {
				kekuatanKenshin -= 6
				lawanDikalahkan++
				break
			} else {
				kecepatanKenshin -= 6
				lawanDikalahkan++
				break
			}
		}
	}

	fmt.Println(lawanDikalahkan, kekuatanKenshin, kecepatanKenshin)

}
