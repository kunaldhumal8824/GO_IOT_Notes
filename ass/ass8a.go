package main

import "fmt"

// Define a structure for Book
type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

// Function to display the details of a book
func (b Book) displayDetails() {
	fmt.Printf("\nBook ID: %d\n", b.BookID)
	fmt.Printf("Title: %s\n", b.Title)
	fmt.Printf("Author: %s\n", b.Author)
	fmt.Printf("Price: %.2f\n", b.Price)
}

func main() {
	var n int

	// Accept the number of books
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	// Create a slice of Book to store book details
	books := make([]Book, n)

	// Loop to accept details for each book
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Book %d:\n", i+1)

		// Accepting book details
		fmt.Print("Enter Book ID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Enter Book Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Enter Book Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Enter Book Price: ")
		fmt.Scan(&books[i].Price)
	}

	// Display the details of all books
	fmt.Println("\nDisplaying Book Details:")
	for _, book := range books {
		book.displayDetails()
	}
}

