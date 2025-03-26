package main

import (
	"fmt"
)

// Function to copy elements from one array to another
func copyArray(src []int, dest []int) {
	// Ensure that both arrays have the same length
	if len(src) != len(dest) {
		fmt.Println("Error: Source and destination arrays must have the same length.")
		return
	}

	// Copy elements from src to dest
	for i := 0; i < len(src); i++ {
		dest[i] = src[i]
	}
}

// Function to display array elements
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	// Initialize source array
	src := []int{10, 20, 30, 40, 50}

	// Initialize destination array with the same length as source
	dest := make([]int, len(src))

	// Print source array
	fmt.Println("Source Array:")
	printArray(src)

	// Call copyArray function to copy elements from src to dest
	copyArray(src, dest)

	// Print destination array after copying
	fmt.Println("Destination Array after copying elements from source:")
	printArray(dest)
}
