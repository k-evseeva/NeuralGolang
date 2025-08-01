package main

import "fmt"

func main_03() {
	// creating pointer
	var a int = 10
	var aPtr *int = &a // referencing
	fmt.Println(&a, aPtr)
	fmt.Println(a, *aPtr) // dereferencing
	// nil pointer
	var nilPtr *int
	fmt.Println("Is nilPtr == nil:", nilPtr == nil)
	// pointer in function
	fmt.Println("a =", a)
	modifyValue(aPtr)
	fmt.Println("a =", a, "(modified)")
	// no arithmetics!
}

func modifyValue(ptr *int) {
	*ptr++
}
