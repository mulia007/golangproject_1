package main

import (
	"fmt"
	"strconv"
)

func findUniqueNumbers(nums []int) []int {
	uniqueNums := make(map[int]int)
	var result []int

	// Menghitung frekuensi kemunculan setiap angka
	for _, num := range nums {
		uniqueNums[num]++
	}

	// Memilih angka yang muncul hanya sekali
	for num, count := range uniqueNums {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9}

	// Mengonversi array angka menjadi array string
	stringNumbers := make([]string, len(numbers))
	for i, num := range numbers {
		stringNumbers[i] = strconv.Itoa(num)
	}

	// Mengonversi array string menjadi array angka
	parsedNumbers := make([]int, len(stringNumbers))
	for i, str := range stringNumbers {
		parsedNum, _ := strconv.Atoi(str)
		parsedNumbers[i] = parsedNum
	}

	unique := findUniqueNumbers(parsedNumbers)
	fmt.Println("Angka-angka yang muncul hanya sekali:", unique)
}
