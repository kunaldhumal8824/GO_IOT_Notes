package main

import "fmt"
func main() {
    // Create a buffered channel with a capacity of 5
    ch := make(chan int, 5)
    
    // Store values in the channel
    ch <- 10
    ch <- 20
    ch <- 30
    ch <- 40
    ch <- 50

    // Print the current capacity and length of the channel
    fmt.Println("Channel Capacity:", cap(ch))
    fmt.Println("Channel Length:", len(ch))

    // Read values from the channel
    fmt.Println("Reading values from channel:")
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    // Print the modified length of the channel after reading values
    fmt.Println("Modified Channel Length:", len(ch))
}
