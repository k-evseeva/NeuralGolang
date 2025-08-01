package main

import "fmt"

func main_02() {
	// factorial
	fmt.Println("5! =", factorial(5))
	fmt.Println("10! =", factorial(10))
	// sum of digits
	num := 12345
	fmt.Printf("Sum of digits for %d is %d\n", num, sumOfDigits(num))
	// fibonacci numbers
	fmt.Println("Fourth fibonacci number:", fibonacci(4))
	fmt.Println("Seventh fibonacci number:", fibonacci(7))
}

func factorial(n int) int {
	// basic case
	if n == 0 {
		return 1
	}
	// recursive case
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// basic case
	if n < 10 {
		return n
	}
	// recursive case
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	// basic case
	if n == 1 || n == 2 {
		return 1
	}
	// resursive case
	return fibonacci(n-1) + fibonacci(n-2)
}
