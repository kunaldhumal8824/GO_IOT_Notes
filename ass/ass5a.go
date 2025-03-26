package main

import (
	"fmt"
	"os"
)

func main() {
	// Name of the file to create
	fileName := "example.txt"

	// Create a new file (or overwrite if the file already exists)
	file, err := os.Create(fileName)
	if err != nil {
		// Handle any errors that occurred during file creation
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	// Write some content to the file
	content := "Hello, this is a sample text file created using Go language."
	_, err = file.WriteString(content)
	if err != nil {
		// Handle any errors that occurred during writing to the file
		fmt.Println("Error writing to the file:", err)
		return
	}

	// Success message
	fmt.Println("File created and content written successfully.")
}
