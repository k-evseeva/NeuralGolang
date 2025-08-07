package main

import (
	"fmt"
)

func main_08() {
	// for as while loop
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}
	// for as infinite loop with break
	sum := 0
	for {
		sum += 10
		fmt.Println("Current sum:", sum)
		if sum == 50 {
			break
		}
	}
	// for as while loop with continue
	num := 1
	for num <= 10 {
		if num%2 != 0 {
			num++
			continue
		}
		fmt.Println("Even number:", num)
		num++
	}
}
