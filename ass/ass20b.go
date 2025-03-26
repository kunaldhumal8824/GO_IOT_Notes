package main

import (
	"fmt"
)

func main() {
	// Create a channel
	ch := make(chan string)

	// Start a goroutine to send data to the channel
	go func() {
		ch <- "Hello, World!"
		ch <- "Goodbye, World!"
		close(ch) // Close the channel
	}()

	// Use a for range loop to receive data from the channel
	for msg := range ch {
		fmt.Println(msg)
	}
}
