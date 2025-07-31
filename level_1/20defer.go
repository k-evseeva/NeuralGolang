package main

import "fmt"

func main_20() {
	process1(10)
}

func process1(arg int) {
	defer fmt.Println("Deferred arg value:", arg)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	arg++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of arg:", arg)
}
