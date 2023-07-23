package main

import "fmt"

func mergeArrays(arr1, arr2 []string) []string {
	merged := append(arr1, arr2...)
	merged = removeDuplicates(merged)
	return merged
}

func removeDuplicates(arr []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, value := range arr {
		if encountered[value] == false {
			encountered[value] = true
			result = append(result, value)
		}
	}

	return result
}

func main() {
	array1 := []string{"kazuya", "jin", "lee"}
	array2 := []string{"kayuza", "feng"}

	mergedArray := mergeArrays(array1, array2)
	fmt.Println("Array hasil penggabungan:", mergedArray)
}
