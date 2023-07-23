package main

import (
	"fmt"
	"strings"
)

func findCommonLetters(word1, word2 string) []string {
	letters1 := strings.Split(word1, "")
	letters2 := strings.Split(word2, "")
	commonLetters := make([]string, 0)

	letterMap := make(map[string]bool)
	for _, letter := range letters1 {
		letterMap[letter] = true
	}

	for _, letter := range letters2 {
		if letterMap[letter] {
			commonLetters = append(commonLetters, letter)
		}
	}

	return commonLetters
}

func main() {
	word1 := "AKA"
	word2 := "AKASHI"

	commonLetters := findCommonLetters(word1, word2)
	fmt.Println("Huruf yang sama:", commonLetters)
}
