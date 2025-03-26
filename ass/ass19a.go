package main

import (
	"fmt"
)

// Function that returns both the sum and difference of two integers
func addSubtract(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	a := 10
	b := 5

	// Call the function and capture the returned values
	sum, diff := addSubtract(a, b)

	// Print the results
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
}
