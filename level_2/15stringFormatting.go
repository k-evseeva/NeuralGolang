package main

import "fmt"

func main_15() {

	// Summary
	num := 42
	num2 := 123456
	fmt.Printf("%05d\n", num)
	fmt.Printf("%05d\n", num2)

	greeting := "Hello"
	fmt.Printf("|%10s|\n", greeting)
	fmt.Printf("|%-10s|\n", greeting)
	fmt.Printf("|%3s|\n", greeting)

	message := "Hello\tWorld!"
	rawMessage := `Hello\tWorld!`
	fmt.Println(message)
	fmt.Println(rawMessage)

	sqlQuery := `SELECT * FROM users WHERE age > 30`
	fmt.Println(sqlQuery)
}
