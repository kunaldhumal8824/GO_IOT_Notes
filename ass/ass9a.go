package main

import "fmt"

// Function to check if the number is a palindrome
func isPalindrome(num int) bool {
	// Store the original number
	originalNum := num
	reversedNum := 0
	var remainder int

	// Reverse the number
	for num != 0 {
		// Get the last digit of the number
		remainder = num % 10
		// Build the reversed number
		reversedNum = reversedNum*10 + remainder
		// Remove the last digit
		num /= 10
	}

	// Check if the original number is equal to the reversed number
	return originalNum == reversedNum
}

func main() {
	var number int

	// Accept the number from the user
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Check if the number is a palindrome
	if isPalindrome(number) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
