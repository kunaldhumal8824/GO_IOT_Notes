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
	content := "Savitribai Phule Pune University\nT.Y. B.C.A. (Science) (Semester-VI) Practical Examination\nBCA 367: DSE V Lab (Programming in GO and IoT)\n"

	// Write the content to the file
	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content appended to file successfully.")
}
