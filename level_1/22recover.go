package main

import "fmt"

func main_22() {
	process3()
	fmt.Println("Returned from process")
}

func process3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Start process")
	panic("Something went wrong!")
	// fmt.Println("End process") - warning (unreachable code)
}
