package main

import (
	"fmt"
	"strconv"
)

func swap(data []int, i int) {
	data[i], data[i + 1] = data[i + 1], data[i]
}

func bubbleSort(data []int) {
	for i := 0 ; i < len(data); i++ {
		for j := 0 ; j < len(data) - i - 1; j++ {
			if data[j] > data[j + 1] {
				swap(data, j)
			}
		}
	}
}

func main() {
	var data = make([]int, 0, 10)
	var input string

	fmt.Println("Enter a sequence of up to 10 integers. Press 'X' to stop...")
	for len(data) < 10 {
		fmt.Print("Enter number: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		if input == "X" {
			break
		}

		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		data = append(data, intInput)
	}
	fmt.Println("Stop listening...")

	fmt.Println("Original slice:", data)

	bubbleSort(data)

	fmt.Println("Sorted slice:", data)
}