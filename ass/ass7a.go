package main

import "fmt"

// Function to calculate the transpose of a matrix
func transposeMatrix(matrix [][]int) [][]int {
	// Get the number of rows and columns of the original matrix
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix to store the transpose
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// Calculate the transpose of the matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

// Function to display a matrix
func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	var rows, cols int

	// Accept the number of rows and columns for the matrix
	fmt.Print("Enter the number of rows and columns for the matrix: ")
	fmt.Scan(&rows, &cols)

	// Create a matrix with user-defined size
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Accept matrix elements from the user
	fmt.Println("Enter the elements of the matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i+1, j+1)
			fmt.Scan(&matrix[i][j])
		}
	}

	// Display the original matrix
	fmt.Println("Original Matrix:")
	displayMatrix(matrix)

	// Calculate and display the transpose of the matrix
	transposed := transposeMatrix(matrix)
	fmt.Println("Transpose of the Matrix:")
	displayMatrix(transposed)
}
