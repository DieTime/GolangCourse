package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var slice = make([]int, 0, 3)
	var input string
	var value int

	// Read first value
	fmt.Print("Enter integer value: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for input != "X" {
		// Convert string to int
		value, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add value to slice
		slice = append(slice, value)

		// Sort slice
		sort.Ints(slice)

		// Print slice
		fmt.Println("Slice:", slice)

		// Read next value
		fmt.Print("\nEnter integer value: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Exit...")
}