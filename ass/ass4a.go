package main

import (
	"fmt"
)

// Recursive function to calculate sum of digits
func sumOfDigits(n int) int {
	// Base case: when the number is reduced to 0
	if n == 0 {
		return 0
	}
	// Recursive case: sum the last digit and the sum of the rest of the digits
	return n%10 + sumOfDigits(n/10)
}

func main() {
	// Input: Asking the user to enter a number
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Calling the recursive function and displaying the result
	result := sumOfDigits(num)
	fmt.Printf("The sum of digits of %d is: %d\n", num, result)
}
