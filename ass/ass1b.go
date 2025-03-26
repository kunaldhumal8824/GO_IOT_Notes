package main

import (
	"fmt"
)

// Define a structure to hold student details
type Student struct {
	RollNo    int
	StudName  string
	Mark1     float64
	Mark2     float64
	Mark3     float64
	Total     float64
	Average   float64
}

// Function to calculate total and average marks
func (s *Student) Calculate() {
	s.Total = s.Mark1 + s.Mark2 + s.Mark3
	s.Average = s.Total / 3
}

func main() {
	var n int
	// Accept the number of students
	fmt.Print("Enter the number of students: ")
	fmt.Scanln(&n)

	// Create a slice to store student details
	students := make([]Student, n)

	// Input student details and calculate total and average marks
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scanln(&students[i].RollNo)
		fmt.Print("Student Name: ")
		fmt.Scanln(&students[i].StudName)
		fmt.Print("Enter Mark1: ")
		fmt.Scanln(&students[i].Mark1)
		fmt.Print("Enter Mark2: ")
		fmt.Scanln(&students[i].Mark2)
		fmt.Print("Enter Mark3: ")
		fmt.Scanln(&students[i].Mark3)

		// Calculate total and average for each student
		students[i].Calculate()
	}

	// Display the student details along with total and average marks
	fmt.Println("\nStudent Details:")
	for i := 0; i < n; i++ {
		fmt.Printf("\nStudent %d:\n", i+1)
		fmt.Printf("Roll No: %d\n", students[i].RollNo)
		fmt.Printf("Name: %s\n", students[i].StudName)
		fmt.Printf("Marks: %.2f, %.2f, %.2f\n", students[i].Mark1, students[i].Mark2, students[i].Mark3)
		fmt.Printf("Total Marks: %.2f\n", students[i].Total)
		fmt.Printf("Average Marks: %.2f\n", students[i].Average)
	}
}

