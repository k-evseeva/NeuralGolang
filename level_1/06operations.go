package main

import (
	"fmt"
	"math"
)

func main_06() {
	var a, b int = 11, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division (integer):", result)
	result = a % b
	fmt.Println("Remainder:", result)

	var a2, b2 int = -11, 3
	result = a2 / b2
	fmt.Println("Division (integer):", result) // not quite right
	result = a2 % b2
	fmt.Println("Remainder:", result) // not quite right too

	var pi float64 = 22 / 7
	fmt.Println("PI (integer):", pi)
	pi = 22 / 7.0 // error, if pi has int
	fmt.Println("PI (float):", pi)

	// overflows
	var maxInt int64 = 9223372036854775807
	fmt.Println("Max for int64:", maxInt)
	maxInt++
	fmt.Println("Overflow:", maxInt)
	var maxUint uint64 = 18446744073709551615
	fmt.Println("Max for uint64:", maxUint)
	maxUint++
	fmt.Println("Overflow:", maxUint)

	// underflows
	var smallFloat float64 = 1.0e-123
	fmt.Println("Small float:", smallFloat)
	smallFloat /= math.MaxFloat64
	fmt.Println("Underflow:", smallFloat)
}
