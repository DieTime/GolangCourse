package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	var starts, ends, contains bool

	fmt.Print("Enter your string: ")

	// Scan line from console
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = strings.ToLower(scanner.Text())

	// Conditions
	starts = strings.HasPrefix(str, "i")
	ends = strings.HasSuffix(str, "n")
	contains = strings.Contains(str, "a")

	// Checking that all conditions are met
	if starts && ends && contains {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
