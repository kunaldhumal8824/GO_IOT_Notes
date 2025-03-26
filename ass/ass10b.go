package main

import (
	"fmt"
)

// Function to generate Fibonacci numbers and send them to the channel
func generateFibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x // Send Fibonacci number to the channel
		x, y = y, x+y // Update x and y for the next Fibonacci number
	}
	close(ch) // Close the channel when done
}

// Function to receive from channel and display Fibonacci numbers
func displayFibonacci(ch chan int) {
	for fib := range ch {
		fmt.Println(fib)
	}
}

func main() {
	// Define how many Fibonacci numbers we want to generate
	n := 10

	// Create a channel for Fibonacci numbers
	ch := make(chan int)

	// Start a goroutine to generate Fibonacci numbers
	go generateFibonacci(n, ch)

	// Start another goroutine to display the Fibonacci numbers
	go displayFibonacci(ch)

	// Wait for the program to finish
	// In this case, we use a simple mechanism to wait for goroutines to finish
	// You could use sync.WaitGroup in a more complex example
	select {} // Block forever, allowing goroutines to run
}
