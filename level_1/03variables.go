package main

import "fmt"

// pet := "Luna" - error (this only inside functions)
var pet = "Luna" // global variable

func main_03() {
	var count int
	var firstName string = "John"
	var lastName = "Lennon"
	age := 84
	var count1, count2 = 1, 2
	fmt.Println(firstName, lastName, age, count)
	fmt.Println(count1, count2)
	fmt.Println(pet)
	var pet = "Elisa"
	fmt.Println(pet)
	hello()
	// Default values
	// Numeric types: 0
	// Boolean types: false
	// String type: ""
	// Pointers, slices, maps, functions and structs: nil
	// fmt.Println(name) - scope related error
}

func hello() {
	name := "Alice" // local variable
	fmt.Println("Hello,", name)
}
