package main

import "fmt"

func printMultiplicationTable(size int) {
	fmt.Printf("Tabel Perkalian %d:\n", size)

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	tableSize := 9 // Ukuran tabel perkalian

	printMultiplicationTable(tableSize)
}
