package main

import "fmt"

func main_19() {
	// simple sum function
	result := sum(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
	fmt.Println("Sum of 0 elements:", sum())
	// sum2 function with string as first argument
	statement2, result2 := sum2("1 + 2 + 3 =", 1, 2, 3)
	fmt.Println(statement2, result2)
	statement2, result2 = sum2("Sum of 0 elements:")
	fmt.Println(statement2, result2)
	// sum3 function with int as first argument
	first1, total1 := sum3(10, 1, 2, 3)
	fmt.Printf("First elem = %d\t Sum = %d\n", first1, total1)
	first2, total2 := sum3(4, 3, 4)
	fmt.Printf("First elem = %d\t Sum = %d\n", first2, total2)
	// unpacking slice
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	first3, total3 := sum3(1, nums...)
	fmt.Printf("First elem = %d\t Sum = %d\n", first3, total3)
	// fmt.Println(sumEnhanced(1, 2, 3))
}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

// func sum2(nums ...int, returnString string) (string, int) {} - error

func sum2(returnString string, nums ...int) (string, int) {
	total := 0
	for _, val := range nums {
		total += val
	}
	return returnString, total
}

func sum3(sequence int, nums ...int) (int, int) {
	total := 0
	for _, val := range nums {
		total += val
	}
	return sequence, total
}

// func sumEnhanced(nums ...int) (string, int) {
// 	stat := []rune{}
// 	total := 0
// 	if len(nums) == 0 {
// 		return "Sum of 0 elements:", total
// 	}
// 	for _, val := range nums {
// 		total += val
// 		stat = append(stat, rune(val), ' ', '+', ' ')
// 	}
// 	stat[len(stat)-2] = '='
// 	stat = append(stat, rune(total))
// 	return string(stat), total
// }
