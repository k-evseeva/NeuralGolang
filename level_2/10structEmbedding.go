package main

import "fmt"

func main_10() {
	emp := employee{
		person: person{name: "John", age: 30},
		empId:  "E001",
		salary: 50000,
	}
	fmt.Println("Name:", emp.name) // emp.person.name
	fmt.Println("Age:", emp.age)   // emp.person.age
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)
	emp.person.introduce()
	emp.introduce()
}

type person struct {
	name string
	age  int
}

type employee struct {
	person // embedded struct (anon field)
	// p person - embedded struct (named field)
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi! I'm %s and I'm %d.\n", p.name, p.age)
}

func (e employee) introduce() { // redefine method
	fmt.Printf("Hi! I'm %s.\n", e.empId)
}
