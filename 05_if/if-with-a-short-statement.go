package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func pow1(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow1(3, 3, 10),
	)
}