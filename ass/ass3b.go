package main

import (
	"fmt"
)

// Define an Employee structure
type Employee struct {
	eno   int
	ename string
	salary float64
}

// Function to find the employee with maximum salary
func findMaxSalaryEmployee(employees []Employee) Employee {
	maxSalaryEmployee := employees[0] // Start by assuming the first employee has the highest salary
	for _, emp := range employees {
		if emp.salary > maxSalaryEmployee.salary {
			maxSalaryEmployee = emp
		}
	}
	return maxSalaryEmployee
}

func main() {
	// Input: Accept the number of employees
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	// Declare a slice to hold employee records
	employees := make([]Employee, n)

	// Input employee records
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for employee %d\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].eno)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].ename)
		fmt.Print("Employee Salary: ")
		fmt.Scan(&employees[i].salary)
	}

	// Find the employee with maximum salary
	maxSalaryEmployee := findMaxSalaryEmployee(employees)

	// Output: Display the employee with the maximum salary
	fmt.Printf("\nEmployee with Maximum Salary:\n")
	fmt.Printf("Employee Number: %d\n", maxSalaryEmployee.eno)
	fmt.Printf("Employee Name: %s\n", maxSalaryEmployee.ename)
	fmt.Printf("Employee Salary: %.2f\n", maxSalaryEmployee.salary)
}
