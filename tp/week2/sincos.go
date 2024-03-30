package main

import (
	"fmt"
	"math"
)

func main() {
	var R, T, lengthX, lengthY float64

	fmt.Scan(&R, &T)

	lengthX = panjangX(R, T)
	lengthY = panjangY(R, T)

	fmt.Printf("%.2f %.2f", lengthX, lengthY)
}

func panjangX(R, T float64) float64 {
	T = T * math.Pi / 180
	return R * math.Cos(T)
}

func panjangY(R, T float64) float64 {
	T = T * math.Pi / 180
	return R * math.Sin(T)
}
