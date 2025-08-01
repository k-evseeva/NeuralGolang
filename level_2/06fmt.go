package main

import "fmt"

func main_06() {

	// printing functions
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print(12, 456)

	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(12, 456)

	name := "Alice"
	age := 18
	fmt.Printf("Name: %s\tAge: %d\n", name, age)
	fmt.Printf("Binary name: %x\tBinary age: %b\n", name, age)

	// formatting functions
	s := fmt.Sprint("Hello,", "World!", 123, 456)
	fmt.Println(s)

	s = fmt.Sprintln("Hello,", "World!", 123, 456)
	fmt.Println(s)

	sf := fmt.Sprintf("|%s|", s)
	fmt.Println(sf)
	sf = fmt.Sprintf("Name: %s\tAge: %d", name, age)
	fmt.Println(sf)

	// scanning functions
	var num1, num2 int
	fmt.Print("Enter two int numbers: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)

	var name2 string
	var num3 int
	fmt.Print("Enter your name and number: ")
	fmt.Scanln(&name2, &num3)
	fmt.Printf("Hello, %s!\n", name2)
	fmt.Printf("Your favorite number is %d.\n", num3)

	var message string
	var user_id int
	fmt.Print("Enter your message and id (separated by space): ")
	fmt.Scanf("%s %d", &message, &user_id)
	fmt.Printf("Message '%s' from user '%d' was sented!\n", message, user_id)

	// error formatting functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You can drive")
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d in too young to drive", age)
	}
	return nil
}
