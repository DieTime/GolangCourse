/*
	It is expected that as the two goroutines
	are executed, there will be alternating
	output similar to this:

	counter: 1
	counter: 2
	counter: 3
	...

	But since the goroutines are not synchronized,
	race conditions arise. And the output turns
	into something like this:

	counter: 2
	counter: 2
	counter: 2
	counter: 4
	...
*/

package main

import (
	"fmt"
	"time"
)

// Global variable for to demonstrate racing
var counter int

// Function for print counter value
func printCounter() {
	fmt.Println("counter:", counter)
}

// Function for increase counter value
func inkCounter() {
	counter += 1
}

func main() {
	// Increase and print counter value 10 times
	for i := 0; i < 10; i++ {
		go inkCounter()
		go printCounter()
	}

	// Waiting for the program not to end
	time.Sleep(5000)
}
