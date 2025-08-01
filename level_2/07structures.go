package main

import "fmt"

func main_07() {
	// structs
	p1 := Person{"John", "Doe", 30, Address{}, PhoneHomeCell{}}
	p2 := Person{firstName: "Jane", age: 25, address: Address{"Berlin", "Germany"}}
	p3 := Person{
		firstName: "Alice",
		lastName:  "Miller ",
		age:       18,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	// accessing fields
	p2.address.city = "New York"
	p2.address.country = "USA"
	fmt.Println(p2.address)
	fmt.Println(p1.firstName, p2.firstName, p3.firstName)
	// anon structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "user@exemple.com",
	}
	fmt.Printf("User: %s\tEmail:%s\n", user.username, user.email)
	// methods
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p3.fullName())
	// method with pointer
	p4 := Person{
		firstName: "Matthew",
		lastName:  "Smith",
		age:       27,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "+7383493984",
			cell: "+7238393292",
		},
	}
	fmt.Println(p4)
	p4.tryIncrementAge()
	fmt.Println(p4)
	p4.incrementAge()
	fmt.Println(p4)
	// embedded struct and anon field
	fmt.Println(p4.address.city) // embedded struct
	fmt.Println(p4.cell)         // anon field
	// compare structs
	fmt.Printf("Is p1 and p2 equal: %t\n", p1 == p2)
	fmt.Printf("p1 != p2: %t\n", p1 != p2)
}

// struct
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address // embedded struct
	PhoneHomeCell         // anon field
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) tryIncrementAge() {
	p.age++
}

func (p *Person) incrementAge() {
	p.age++
}

// another structs
type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}
