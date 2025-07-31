package main

import "fmt"

func main_17() {
	// simple function
	sum := add(1, 2)
	fmt.Println("1 + 2 =", sum)
	fmt.Println("2 + 3 =", add(2, 3))
	// anonymous function
	greet := func() {
		fmt.Println("Hello, anonymous function!")
	}
	greet()
	// function as object
	operation := add
	result := operation(3, 5)
	fmt.Println("3 + 5 =", result)
	// passing function as argument
	result = applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)
	// returning and using function
	multBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multBy2(6))
}

// simple function
func add(a int, b int) int {
	return a + b
}

// function that takes function as argument
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
