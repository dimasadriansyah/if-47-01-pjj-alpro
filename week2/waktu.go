package main

import "fmt"

func main() {
	var hours, minutes, seconds int
	var total int = 0

	fmt.Scan(&hours, &minutes, &seconds)

	waktu(hours, minutes, seconds, &total)
	fmt.Println("Hasil konversi =", total, "detik")

}

func waktu(jam, menit, detik int, total *int) {

	*total = *total + (jam*60)*60
	*total = *total + menit*60
	*total = *total + detik
}
