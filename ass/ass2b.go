package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the file path (replace with your file path)
	filePath := "example.txt"

	// Get file information
	fileInfo, err := os.Stat(filePath)

	// Check if there is an error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file information
	fmt.Println("File Name:", fileInfo.Name())              // File Name
	fmt.Println("File Size:", fileInfo.Size(), "bytes")     // File Size in bytes
	fmt.Println("Permissions:", fileInfo.Mode())           // File Permissions
	fmt.Println("Last Modified:", fileInfo.ModTime())       // Last Modification Time
	fmt.Println("Is Directory:", fileInfo.IsDir())          // Check if it's a directory
}
