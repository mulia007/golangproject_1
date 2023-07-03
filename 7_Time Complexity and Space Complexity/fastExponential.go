package main

import "fmt"

func power(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		temp := power(x, n/2)
		return temp * temp
	}

	temp := power(x, (n-1)/2)
	return x * temp * temp
}

func main() {
	x := 2  // Basis
	n := 10 // Eksponen

	result := power(x, n)
	fmt.Println(x, "^", n, "=", result)
}
