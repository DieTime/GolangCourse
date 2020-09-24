package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	name    string
	surname string
}

func main() {
	var filename string
	var slice = make([]Person, 0)

	// Get filename
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter path to file: ")
	scanner.Scan()
	filename = scanner.Text()

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read file line by line
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		// Create empty struct
		person := Person{}

		// Read first name and last name from line to struct
		n, err := fmt.Sscanf(scanner.Text(), "%s %s", &person.name, &person.surname)
		if n != 2 || err != nil {
			fmt.Println("Warning! A read error has occurred!")
		}

		// Add struct to slice
		slice = append(slice, person)
	}

	// Print slice
	for index, item := range slice {
		fmt.Println("Person", index + 1, ":", item.name, item.surname)
	}

	// Close file
	err = file.Close()
	if err != nil {
		fmt.Println("Warning!", err)
	}
}