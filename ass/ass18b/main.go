package main

import (
	"ass18b/calculator" // Replace with your actual module name
	"fmt"
)

func main() {
	var a, b int
	var choice int

	fmt.Println("Simple Calculator")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Print("Choose an operation (1-4): ")
	fmt.Scan(&choice)

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	var result float64
	switch choice {
	case 1:
		result = float64(calculator.Add(a, b))
	case 2:
		result = float64(calculator.Subtract(a, b))
	case 3:
		result = float64(calculator.Multiply(a, b))
	case 4:
		result = calculator.Divide(a, b)
	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
