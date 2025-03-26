package main

import (
	"fmt"
)

type Employee struct {
	eno    int
	ename  string
	salary float64
}

// Function to find and display employees with the minimum salary
func displayMinSalaryEmployees(employees []Employee) {
	if len(employees) == 0 {
		fmt.Println("No employee records available.")
		return
	}

	// Find the minimum salary
	minSalary := employees[0].salary
	for _, emp := range employees {
		if emp.salary < minSalary {
			minSalary = emp.salary
		}
	}

	// Display the employees having the minimum salary
	fmt.Println("\nEmployees with the minimum salary:")
	for _, emp := range employees {
		if emp.salary == minSalary {
			fmt.Printf("Employee Number: %d, Name: %s, Salary: %.2f\n", emp.eno, emp.ename, emp.salary)
		}
	}
}

func main() {
	var n int

	// Accept number of employees
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	// Slice to store employee records
	employees := make([]Employee, n)

	// Accept employee information
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for employee %d:\n", i+1)
		fmt.Print("Enter employee number: ")
		fmt.Scan(&employees[i].eno)
		fmt.Print("Enter employee name: ")
		fmt.Scan(&employees[i].ename)
		fmt.Print("Enter employee salary: ")
		fmt.Scan(&employees[i].salary)
	}

	// Call the function to display employees with the minimum salary
	displayMinSalaryEmployees(employees)
}
