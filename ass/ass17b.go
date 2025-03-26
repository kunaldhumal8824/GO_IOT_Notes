package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Content to append
	content := "This is a new line of text.\n"

	// Write the content to the file
	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content appended to file successfully.")
}
