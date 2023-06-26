package main

import "fmt"

func power(base, exponent float64) float64 {
	result := 1.0
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	return result
}

func main() {
	x := 2.0
	y := 3.0

	result := power(x, y)
	fmt.Printf("%.2f^%.2f = %.2f\n", x, y, result)
}
