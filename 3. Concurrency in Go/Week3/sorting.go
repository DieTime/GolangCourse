package main

import (
	"fmt"
	"strconv"
)

// Function for slice sorting
func Sort(block []int, channel chan []int) {
	// Sort block
	for i := 0; i < len(block); i++ {
		for j := 0; j < len(block)-i-1; j++ {
			if block[j] > block[j+1] {
				block[j], block[j+1] = block[j+1], block[j]
			}
		}
	}

	// Send result to channel
	channel <- block
}

// Function for merging 2 sorted slices
func Merge(ldata []int, rdata []int, channel chan []int) {
	// Create result slice
	result := make([]int, 0, len(ldata)+len(rdata))

	// Left and right cursors
	l, r := 0, 0

	// Merging slices
	for i := 0; i < cap(result); i++ {
		switch {
		case l >= len(ldata):
			result = append(result, rdata[r])
			r += 1
		case r >= len(rdata):
			result = append(result, ldata[l])
			l += 1
		case ldata[l] < rdata[r]:
			result = append(result, ldata[l])
			l += 1
		default:
			result = append(result, rdata[r])
			r += 1
		}
	}

	// Send result to channel
	channel <- result
}

func main() {
	var data = make([]int, 0, 10)
	var channel = make(chan []int, 4)
	var input string

	// Reading integers
	fmt.Println("Enter a sequence of integers. Press 'X' to stop...")
	for {
		fmt.Print("Enter number: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Stop reading if input is 'X'
		if input == "X" {
			break
		}

		// Convert input to integer
		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add integer to slice
		data = append(data, intInput)
	}
	fmt.Println("Stop listening...")

	// Calculating block size
	blockSize := len(data) / 4

	// Start sorting 4 blocks concurrently
	for i := 0; i < 4; i++ {

		/* Calculate begin and
		   end of the block  */

		start := i * blockSize
		end := (i + 1) * blockSize

		if i == 3 {
			end = len(data)
		}

		// Start sorting of the block
		go Sort(data[start:end], channel)
	}

	/* Waiting for blocks sorting
	   and processing of results */

	a := <-channel
	b := <-channel
	c := <-channel
	d := <-channel

	// Pairing blocks a, b
	go Merge(a, b, channel)

	// Pairing blocks c, d
	go Merge(c, d, channel)

	/* Waiting for blocks merging
	   and processing of results */

	ab := <-channel
	cd := <-channel

	// Pairing 2 merged blocks ab and cd
	go Merge(ab, cd, channel)

	// Waiting result slice
	result := <-channel

	// Printing result slice
	fmt.Println("Sorted array:", result)
}