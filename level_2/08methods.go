package main

import "fmt"

func main_08() {

	rect := Rectangle{10, 9}
	area := rect.Area()
	fmt.Printf("Area of rect = %.2f\n", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Printf("Area of scaled rect = %.2f\n\n", area)

	num1 := MyInt(-5)
	num2 := MyInt(9)
	fmt.Println(num1.welcomeMessage())
	fmt.Printf("Is num1=%d positive: %t\n", num1, num1.IsPositive())
	fmt.Printf("Is num2=%d positive: %t\n\n", num2, num2.IsPositive())

	sh := Shape{Rectangle: Rectangle{7, 5}}
	fmt.Println("New rectangle area:")
	fmt.Println(sh.Rectangle.Area())
	fmt.Println(sh.Area())
}

// Struct
type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// Define new type
type MyInt int

// Methods on user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt type!"
}

// struct with embedded struct (anon field)
type Shape struct {
	Rectangle
}
