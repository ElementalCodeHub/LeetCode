package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	// Reverse the digits of the number
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

func main() {
	fmt.Println(isPalindrome(121))
}
