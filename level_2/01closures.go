package main

import "fmt"

func main_01() {
	sequence := adder()
	fmt.Println("Current value of i:", sequence())
	fmt.Println("Current value of i:", sequence())
	fmt.Println("Current value of i:", sequence())
	sequence2 := adder()
	fmt.Println("Current value of i:", sequence2())
	fmt.Println("Current value of i:", sequence())
	subtracter := func() func(int) int {
		countDown := 99
		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))
}

func adder() func() int {
	i := 0
	fmt.Println("Init value of i:", i)
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
