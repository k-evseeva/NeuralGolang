package main

import (
	"fmt"
	"slices"
)

func main_14() {
	// creating slices
	var nums1 []int // nil slice
	var nums2 = []int{1, 2, 3}
	nums3 := []int{1, 2, 3, 4, 5}
	// len and cap
	fmt.Printf("nums1: %v\t\tlen: %d\tcap:%d\n", nums1, len(nums1), cap(nums1))
	fmt.Printf("nums2: %v\t\tlen: %d\tcap:%d\n", nums2, len(nums2), cap(nums2))
	fmt.Printf("nums3: %v\tlen: %d\tcap:%d\n", nums3, len(nums3), cap(nums3))
	// creating with make
	slc1 := make([]int, 5)
	slc2 := make([]int, 5, 8)
	// len and cap
	fmt.Println("slc1:", slc1)
	fmt.Println("Length of slc1:", len(slc1))
	fmt.Println("Capacity of slc1:", cap(slc1))
	fmt.Println("slc2:", slc2)
	fmt.Println("Length of slc2:", len(slc2))
	fmt.Println("Capacity of slc2:", cap(slc2))
	// creating with base array
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println("array:", array)
	fmt.Println("slice based on array:", slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice)) // associated with base array
	// creating slice based on slice
	slice2 := slice[2:4]
	fmt.Println("slice2:", slice2)
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))
	// accessing by index
	fmt.Println("Third element:", slice[2])
	slice[2] = 15
	fmt.Println("Third element:", slice[2])
	// append elements
	slice = append(slice, 10, 11)
	// copy slice
	copiedSlice := make([]int, len((slice)))
	copy(copiedSlice, slice)
	fmt.Println("Original slice:", slice)
	fmt.Println("Copied slice:", copiedSlice)
	copiedSlice[2] = 100
	fmt.Println("Original slice:", slice)
	fmt.Println("Copied slice (changed):", copiedSlice)
	slice[1] = -7
	fmt.Println("Original slice (changed):", slice)
	fmt.Println("Copied slice:", copiedSlice)
	// iteration over slice
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	// comparing slices
	if slices.Equal(slice, copiedSlice) {
		fmt.Println("slice is equal to copiedSlice")
	} else {
		fmt.Println("slice is not equal to copiedSlice")
		fmt.Println("slice:", slice)
		fmt.Println("copiedSlice:", copiedSlice)
	}
	// two-dimensional slice
	matrix := make([][]int, 4)
	for i := 0; i < len(matrix); i++ {
		rowLen := i + 1
		matrix[i] = make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			matrix[i][j] = i + j
			fmt.Printf("Add value %d at indices %d %d\n", i+j, i, j)
		}
	}
	fmt.Println(matrix)
}
