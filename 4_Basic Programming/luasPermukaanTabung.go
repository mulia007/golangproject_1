package main

import (
	"fmt"
	"math"
)

func main() {
	var T, r float64
	fmt.Print("Masukkan tinggi tabung (T): ")
	fmt.Scanln(&T)
	fmt.Print("Masukkan jari-jari tabung (r): ")
	fmt.Scanln(&r)

	luasPermukaan := 2 * math.Pi * r * (r + T)
	fmt.Printf("Luas permukaan tabung adalah: %.3f\n", luasPermukaan)
}
