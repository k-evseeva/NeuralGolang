package main

import "fmt"

func main_13() {
	// switch statement
	fruit := "pineapple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown fruit")
	}
	// multiple conditions
	day := "Tue"
	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("It's a weekday")
	case "Sat", "Sun":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}
	// special switch
	num := 15
	switch {
	case num < 10:
		fmt.Println("Number is less than 10")
	case 10 <= num && num < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Unknown number")
	}
	// fallthrough
	x := 2
	switch {
	case x > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case x == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not 2")
	}
	// check type
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
	// if
	scoreIf := 85
	if scoreIf >= 90 {
		fmt.Println("Grade A")
	} else if scoreIf >= 80 {
		fmt.Println("Grade B")
	} else if scoreIf >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
	// switch
	scoreSwitch := 85
	switch {
	case scoreSwitch >= 90:
		fmt.Println("Grade A")
	case scoreSwitch >= 80:
		fmt.Println("Grade B")
	case scoreSwitch >= 70:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade D")
	}
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
		// fallthrough - error
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
