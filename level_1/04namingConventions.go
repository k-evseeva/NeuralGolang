package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main_04() {
	fmt.Println("PascalCase")
	fmt.Println("For structs, interfaces, enums")
	fmt.Println("E.g. CalculateArea, UserInfo, NewHTTPRequest")
	fmt.Println()
	fmt.Println("snake_case")
	fmt.Println("For variables, constants, file names")
	fmt.Println("E.g. user_id, first_name, http_request")
	fmt.Println()
	fmt.Println("UPPERCASE")
	fmt.Println("For constants")
	fmt.Println("E.g. PI, E")
	fmt.Println()
	fmt.Println("mixedCase")
	// either snake_case or mixedCase
	fmt.Println("For variables, constants, file names")
	fmt.Println("E.g. javaScript, htmlDocument, isValid")
	// practice
	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID:", employeeID)
}
