package main

import "fmt"

func main() {

	var s1, s2, s3 float64
	var segitigaSS bool
	fmt.Scan(&s1, &s2, &s3)
	segitigaSS = segitigaSamaSisi(s1, s2, s3)
	fmt.Println(segitigaSS)

}

func segitigaSamaSisi(a, b, c float64) bool {

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}

	return a == b && b == c
}
