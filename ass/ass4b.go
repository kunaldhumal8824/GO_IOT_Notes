package main

import (
	"fmt"
)

// Bubble sort function to sort an array in ascending order
func bubbleSort(arr []int) {
	n := len(arr)
	// Loop over each element in the array
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted, so we reduce the range
		for j := 0; j < n-i-1; j++ {
			// Swap if the element is greater than the next element
			if arr[j] > arr[j+1] {
				// Swap the elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Input: Asking the user to enter array elements
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	// Creating an array of n elements
	arr := make([]int, n)

	// Taking input for array elements
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Sorting the array using bubble sort
	bubbleSort(arr)

	// Output: Displaying the sorted array
	fmt.Println("Array elements in ascending order:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
