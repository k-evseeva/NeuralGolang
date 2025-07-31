package main

import "fmt"

func main_16() {
	// print message runes (and their codes)
	message := "Hello, Go!"
	for ind, val := range message {
		fmt.Println("Index and code:", ind, val)
		fmt.Printf("Index and rune: %d %c\n", ind, val)
	}
	// trying to change slice
	slice := []int{1, 2, 3}
	fmt.Println("slice:", slice)
	for _, val := range slice {
		_ = val
		val = 7 // it doesn't change slice (copy of element)
	}
	fmt.Println("slice (not changed):", slice)
	for idx := range slice {
		slice[idx] = 7
	}
	fmt.Println("slice (changed):", slice)
}
