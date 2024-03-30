package main

import "fmt"

func main() {
	var i, cAwal, cAkhir, step, C, R, F, K float64

	fmt.Scan(&cAwal, &cAkhir, &step)

	fmt.Printf("%10s %10s %10s %10s \n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")

	for i = cAwal; i <= cAkhir; {
		C = i
		R = reamur(i)
		F = fahrenheit(i)
		K = kelvin(i)

		fmt.Printf("%10.2f %10.2f %10.2f %10.2f \n", C, R, F, K)

		i = i + step
	}

}

func reamur(C float64) float64 {
	return C * 0.8
}

func fahrenheit(C float64) float64 {
	return (C * 9 / 5) + 32
}

func kelvin(C float64) float64 {
	return C + 273.00
}
