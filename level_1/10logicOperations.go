package main

import "fmt"

func main_10() {
	var a, b bool = false, false
	var c, d bool = true, true
	// logical NOT
	fmt.Println("NOT false:", !a)
	fmt.Println("NOT true:", !c)
	// logical OR
	fmt.Println("false OR false:", a || b)
	fmt.Println("false OR true:", a || c)
	fmt.Println("true OR false:", c || a)
	fmt.Println("true OR true:", c || d)
	fmt.Println("false OR true OR false", a, c, b)
	// logical AND
	fmt.Println("false AND false:", a && b)
	fmt.Println("false AND true:", a && c)
	fmt.Println("true AND false:", c && a)
	fmt.Println("true AND true:", c && d)
	fmt.Println("true AND false AND true", c, a, d)
	// bitwise AND (&)
	// bitwise OR (|)
	// bitwise XOR (^)
	// bitwise AND NOT (&^)
	// left shift (<<)
	// right shift (>>)
	var x, y, z int = 5, 8, 5
	// equal
	fmt.Println("Equal 5 and 8:", x == y)
	fmt.Println("Equal 5 and 5:", x == z)
	// not equal
	fmt.Println("Not equal 5 and 8:", x != y)
	fmt.Println("Not equal 5 and 5:", x != z)
	// less than
	fmt.Println("5 less than 8:", x < y)
	fmt.Println("8 less than 5:", y < x)
	fmt.Println("5 less than 5:", x < z)
	// less than or equal to
	fmt.Println("5 less or equal than 8:", x <= y)
	fmt.Println("8 less or equal than 5:", y <= x)
	fmt.Println("5 less or equal than 5:", x <= z)
	// greater than
	fmt.Println("5 greater than 8:", x > y)
	fmt.Println("8 greater than 5:", y > x)
	fmt.Println("5 greater than 5:", x > z)
	// greater that or equal to
	fmt.Println("5 greater or equal than 8:", x >= y)
	fmt.Println("8 greater or equal than 5:", y >= x)
	fmt.Println("5 greater or equal than 5:", x >= z)
}
