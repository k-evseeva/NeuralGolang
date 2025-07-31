package main

import "fmt"

func main_12() {
	// creating arrays
	var nums [5]int
	fmt.Println("Numbers:", nums)
	fruits := [4]string{"apple", "banana", "orange", "grape"}
	fmt.Println("Fruits:", fruits)
	// accessing by index
	nums[0] = 9
	nums[2] = 10
	nums[4] = 11
	fmt.Println(nums)
	fmt.Println("Third fruit:", fruits[2])
	// length of array
	fmt.Println("The length of nums array is", len(nums))
	// copying array
	origArray := [3]int{1, 2, 3}
	copiedArray := origArray
	copiedArray[0] = 100
	fmt.Println("Original array:", origArray)
	fmt.Println("Copied array:", copiedArray)
	// change array
	var sameArray *[3]int = &origArray
	sameArray[0] = 100
	fmt.Println("Original array:", origArray)
	fmt.Println("Same array:", *sameArray)
	// iteration over array
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Element at index %d: %d\n", i, nums[i])
	}
	for index, value := range nums {
		fmt.Printf("Index: %d\tValue: %d\n", index, value)
	}
	// iteration with _ (blank identifiers)
	for _, value := range nums {
		fmt.Printf("Value: %d\n", value)
	}
	for index := range nums { // for index, _ := range nums
		fmt.Printf("Index: %d\n", index)
	}
	// about blank identifiers
	a, _ := someFunction()
	fmt.Println(a)
	_, b := someFunction()
	fmt.Println(b)
	c := 3
	_ = c // now there is no error "declared and not used"
	// comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{1, 2, 4}
	// array4 := [2]int{1, 2}
	fmt.Println("array1 == array2:", array1 == array2)
	fmt.Println("array1 == array3:", array1 == array3)
	// fmt.Println("array1 == array4:", array1 == array4) - error "mismatched types"
	// two-dimensional arrays (aka matrices)
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

}

func someFunction() (int, int) {
	return 1, 2
}
