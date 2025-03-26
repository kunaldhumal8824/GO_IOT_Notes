package main

import "fmt"

// Define the Student structure
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// Define a method 'show' with a pointer receiver
func (s *Student) show() {
	fmt.Printf("Student ID: %d\n", s.ID)
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Student Age: %d\n", s.Age)
	fmt.Printf("Student Grade: %s\n", s.Grade)
}

func main() {
	// Create a new Student instance
	student1 := Student{
		ID:    101,
		Name:  "John Doe",
		Age:   20,
		Grade: "A",
	}

	// Call the 'show' method with the pointer receiver
	student1.show()
}
