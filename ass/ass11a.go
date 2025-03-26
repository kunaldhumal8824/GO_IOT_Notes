package main

import "fmt"

func main() {
    // Declare a variable to hold the number
    var num int
    
    // Ask the user for input
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // Check if the number is a two-digit number
    if num >= 10 && num <= 99 || num <= -10 && num >= -99 {
        fmt.Println("The number is a two-digit number.")
    } else {
        fmt.Println("The number is not a two-digit number.")
    }
}

