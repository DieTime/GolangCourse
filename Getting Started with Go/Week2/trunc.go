package main

import "fmt"

func main() {
	var floatInput float32
	var intInput int

	fmt.Print("Enter your float value: ")

	_, err := fmt.Scan(&floatInput)
	if err != nil {
		fmt.Println(err)
	}

	intInput = int(floatInput)

	fmt.Println("Truncated value:", intInput)
}