package main

import "fmt"

func main_07() {
	// simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// iteration over collection
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}
	// break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", i)
		if i == 5 {
			break
		}
	}
	// nested loops
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // new line
	}
	// new feature
	for i := range 5 {
		fmt.Println(5 - i)
	}
	for i := range 5 {
		i++
		fmt.Println(i)
	}
}
