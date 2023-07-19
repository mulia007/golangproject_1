package main

import "fmt"

func findMinMax(numbers []int, min, max *int) {
	for _, num := range numbers {
		if num < *min {
			*min = num
		}
		if num > *max {
			*max = num
		}
	}
}

func main() {
	numbers := []int{1, 2, 3, 9, 7, 8}

	// Initialize min and max with the first element of the slice
	min := numbers[0]
	max := numbers[0]

	findMinMax(numbers, &min, &max)

	fmt.Println("Minimum value:", min)
	fmt.Println("Maximum value:", max)
}
