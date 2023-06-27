package main

import (
	"fmt"
	"math"
)

func main() {
	var r, t float64
	fmt.Print("jari-jari tabung: 4. ")
	fmt.Scanln(&r)
	fmt.Print("tinggi tabung: 20. ")
	fmt.Scanln(&t)

	luasAlas := math.Pi * math.Pow(r, 2)
	luasSelimut := 2 * math.Pi * 4 * 20
	luasPermukaan := luasAlas + luasSelimut

	fmt.Println("Luas permukaan tabung adalah:", luasPermukaan)
}
