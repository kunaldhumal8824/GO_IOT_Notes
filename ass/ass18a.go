package main

import (
	"fmt"
)

// Function to print the multiplication table of a given number
func printMultiplicationTable(number int) {
	fmt.Printf("Multiplication Table of %d:\n", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

func main() {
	var number int
	fmt.Print("Enter a number to print its multiplication table: ")
	fmt.Scan(&number)
	printMultiplicationTable(number)
}
