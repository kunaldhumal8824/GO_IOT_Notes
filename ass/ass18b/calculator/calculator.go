package calculator

// Add function to add two numbers
func Add(a, b int) int {
	return a + b
}

// Subtract function to subtract two numbers
func Subtract(a, b int) int {
	return a - b
}

// Multiply function to multiply two numbers
func Multiply(a, b int) int {
	return a * b
}

// Divide function to divide two numbers
func Divide(a, b int) float64 {
	if b == 0 {
		return 0 // Handle division by zero
	}
	return float64(a) / float64(b)
}
