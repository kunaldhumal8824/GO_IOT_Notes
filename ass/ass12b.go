package main

import (
    "fmt"
    "sync"
)

// Function to check if a number is even or odd and send it to the respective channel
func checkEvenOdd(num int, evenChan chan int, oddChan chan int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement the counter when the goroutine completes
    if num%2 == 0 {
        evenChan <- num
    } else {
        oddChan <- num
    }
}

func main() {
    // Slice of integers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Channels for even and odd numbers
    evenChan := make(chan int)
    oddChan := make(chan int)

    // WaitGroup to wait for all goroutines to finish
    var wg sync.WaitGroup

    // Start the goroutines to check even/odd for each number
    for _, num := range numbers {
        wg.Add(1)
        go checkEvenOdd(num, evenChan, oddChan, &wg)
    }

    // Close channels once all goroutines are finished
    go func() {
        wg.Wait()
        close(evenChan)
        close(oddChan)
    }()

    // Display results from even and odd channels
    fmt.Println("Even Numbers:")
    for num := range evenChan {
        fmt.Println(num)
    }

    fmt.Println("Odd Numbers:")
    for num := range oddChan {
        fmt.Println(num)
    }
}
