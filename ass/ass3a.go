package main

import (
	"fmt"
	"strconv"
)

// Function to check if a number is a palindrome
func isPalindrome(number int) bool {
	// Convert the number to a string
	strNumber := strconv.Itoa(number)
	// Get the length of the string
	length := len(strNumber)

	// Compare the characters from the start and end
	for i := 0; i < length/2; i++ {
		if strNumber[i] != strNumber[length-i-1] {
			return false // If characters do not match, return false
		}
	}
	return true // If the loop completes, the number is a palindrome
}

func main() {
	// Input: Asking the user to enter a number
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is a palindrome using the isPalindrome function
	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome number.")
	} else {
		fmt.Println(num, "is not a palindrome number.")
	}
}
