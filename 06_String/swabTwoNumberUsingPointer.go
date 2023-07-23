package main

import "fmt"

func swapNumbers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	num1 := 10
	num2 := 20

	fmt.Println("Before swapping:")
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

	// Pass the addresses of num1 and num2 to the swapNumbers function
	swapNumbers(&num1, &num2)

	fmt.Println("After swapping:")
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)
}
