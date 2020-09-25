package main

import "fmt"

// Generate function to calculate displacement at time
func GenDisplaceFn(s0, v0, a float64) func (t float64) float64 {
	return func(t float64) float64 {
		return s0 + v0 * t + a * t * t / 2
	}
}

func main() {
	var s0, v0, a, t float64

	// Read start parameters
	fmt.Print("Please enter s0, v0 and a through a space: ")
	_, err := fmt.Scanln(&s0, &v0, &a)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate functions based on start parameters
	DisplaceFn := GenDisplaceFn(s0, v0, a)

	// Read time
	fmt.Print("Please enter t: ")
	_, err = fmt.Scanln(&t)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result
	fmt.Println("Your displacement at time t:", DisplaceFn(t))
}
