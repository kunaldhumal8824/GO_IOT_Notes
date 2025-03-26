package main

import (
	"fmt"
)

// Function to multiply two matrices
func multiplyMatrices(matrixA, matrixB [][]int) [][]int {
	// Get dimensions of matrices
	rowsA := len(matrixA)
	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	// Ensure matrices can be multiplied (cols of A must equal rows of B)
	if colsA != rowsB {
		fmt.Println("Error: Matrices cannot be multiplied!")
		return nil
	}

	// Create result matrix initialized to zero
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// Perform multiplication
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

// Function to print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	var rowsA, colsA, rowsB, colsB int

	// Input dimensions of matrix A
	fmt.Print("Enter number of rows and columns for Matrix A: ")
	fmt.Scan(&rowsA, &colsA)

	// Input dimensions of matrix B
	fmt.Print("Enter number of rows and columns for Matrix B: ")
	fmt.Scan(&rowsB, &colsB)

	// Matrix A and Matrix B dimensions must be compatible for multiplication
	if colsA != rowsB {
		fmt.Println("Error: Matrices cannot be multiplied! The number of columns of Matrix A must equal the number of rows of Matrix B.")
		return
	}

	// Declare matrices
	matrixA := make([][]int, rowsA)
	matrixB := make([][]int, rowsB)

	// Input Matrix A
	fmt.Println("Enter elements of Matrix A:")
	for i := 0; i < rowsA; i++ {
		matrixA[i] = make([]int, colsA)
		for j := 0; j < colsA; j++ {
			fmt.Printf("Enter element at position (%d,%d): ", i+1, j+1)
			fmt.Scan(&matrixA[i][j])
		}
	}

	// Input Matrix B
	fmt.Println("Enter elements of Matrix B:")
	for i := 0; i < rowsB; i++ {
		matrixB[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			fmt.Printf("Enter element at position (%d,%d): ", i+1, j+1)
			fmt.Scan(&matrixB[i][j])
		}
	}

	// Perform matrix multiplication
	result := multiplyMatrices(matrixA, matrixB)

	// Display result
	if result != nil {
		fmt.Println("\nResultant Matrix after multiplication:")
		printMatrix(result)
	}
}
