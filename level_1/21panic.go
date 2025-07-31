package main

import "fmt"

func main_21() {
	// panic(interface{})
	// valid input
	process2(10)
	process2(15)
	process2(25)
	// invalid input
	process2(-1)
	// valid input
	process2(7)
	process2(10)
	process2(13)
}

func process2(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be a non-negative number")
		// defer fmt.Println("After panic") - warning
	}
	fmt.Println("Processing input:", input)
}
