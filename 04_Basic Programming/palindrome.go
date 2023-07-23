package main

import "fmt"

func isPalindrome(word string) bool {
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	word := "katak" // Kata yang ingin diperiksa

	if isPalindrome(word) {
		fmt.Printf("%s adalah palindrome.\n", word)
	} else {
		fmt.Printf("%s bukan palindrome.\n", word)
	}
}
