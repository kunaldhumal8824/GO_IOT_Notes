package main

import "fmt"

// Define an interface called 'Value' with a method to get the value
type Value interface {
	GetValue() interface{}
}

// Define a struct 'Integer' to hold an integer
type Integer struct {
	Value int
}

// Define a struct 'Float' to hold a floating-point number
type Float struct {
	Value float64
}

// Define a struct 'String' to hold a string value
type String struct {
	Value string
}

// Implement the GetValue() method for Integer
func (i Integer) GetValue() interface{} {
	return i.Value
}

// Implement the GetValue() method for Float
func (f Float) GetValue() interface{} {
	return f.Value
}

// Implement the GetValue() method for String
func (s String) GetValue() interface{} {
	return s.Value
}

// Function to display the value using type assertion
func displayValue(v Value) {
	// Use type assertion to check the underlying type and value
	switch v := v.(type) {
	case Integer:
		fmt.Printf("Integer Value: %d\n", v.GetValue())
	case Float:
		fmt.Printf("Float Value: %.2f\n", v.GetValue())
	case String:
		fmt.Printf("String Value: %s\n", v.GetValue())
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// Create instances of Integer, Float, and String
	intValue := Integer{Value: 42}
	floatValue := Float{Value: 3.14159}
	stringValue := String{Value: "Hello, Go!"}

	// Display the values using type assertion
	displayValue(intValue)
	displayValue(floatValue)
	displayValue(stringValue)
}
