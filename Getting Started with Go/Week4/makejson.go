package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Initialize map
	var m map[string]string
	m = make(map[string]string)

	// Initialize scanner for reading from console
	scanner := bufio.NewScanner(os.Stdin)

	// Read name
	fmt.Print("Enter your name: ")
	scanner.Scan()
	m["name"] = scanner.Text()

	// Read address
	fmt.Print("Enter your address: ")
	scanner.Scan()
	m["address"] = scanner.Text()

	// Pack map to json object
	jsonString, err := json.Marshal(m)

	// Handle probably error
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result json
	fmt.Println("Your data in json format:", string(jsonString))
}