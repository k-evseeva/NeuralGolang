package main

import (
	"errors"
	"fmt"
)

func main_18() {
	// simple functions
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v, remainder: %v\n", q, r)
	q, r = divideBy5(43)
	fmt.Printf("Quotient: %v, remainder: %v\n", q, r)
	// fmt.Printf("Quotient: %v, remainder: %v\n", divide(10, 3)) - error
	// checking for error
	result, err := compare(3, 2)
	check(result, err)
	result, err = compare(3, 3)
	check(result, err)
}

// simple function with multiple return
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// named return values
func divideBy5(a int) (quotient int, remainder int) {
	quotient = a / 5
	remainder = a % 5
	return
}

// function that checks error
func compare(a, b int) (string, error) {
	if a > b {
		return "first is greater than second", nil
	} else if b > a {
		return "second is greater than first", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

func check(res string, err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}
