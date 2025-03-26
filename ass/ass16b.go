package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(i*25) * time.Millisecond) // Delay between 0 and 250 ms
	}
}
