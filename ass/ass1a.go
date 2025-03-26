package main

import (
	"fmt"
)

func main() {
	var choice int
	var num1, num2 float64

	// Display the menu to the user
	fmt.Println("Choose an arithmetic operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Enter your choice (1/2/3/4): ")
	fmt.Scanln(&choice)

	// Accept two numbers from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Perform the chosen operation
	switch choice {
	case 1:
		// Addition
		result := num1 + num2
		fmt.Printf("The result of %.2f + %.2f = %.2f\n", num1, num2, result)
	case 2:
		// Subtraction
		result := num1 - num2
		fmt.Printf("The result of %.2f - %.2f = %.2f\n", num1, num2, result)
	case 3:
		// Multiplication
		result := num1 * num2
		fmt.Printf("The result of %.2f * %.2f = %.2f\n", num1, num2, result)
	case 4:
		// Division
		if num2 != 0 {
			result := num1 / num2
			fmt.Printf("The result of %.2f / %.2f = %.2f\n", num1, num2, result)
		} else {
			fmt.Println("Error! Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice! Please choose a valid operation (1-4).")
	}
}
