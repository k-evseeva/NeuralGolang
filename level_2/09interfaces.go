package main

import (
	"fmt"
	"math"
)

func main_09() {
	r := rectangle{3, 4}
	c := circle{5}
	measure(r)
	measure(c)
	fmt.Printf("diameter = %f\n", c.diameter())
	// r2 := rectangle2{4, 5}
	// measure(r2) - error (missing perim)
	myPrinter(1, 2, "John", 48.3, true)
	printType(9)
	printType(3.5)
	printType("Alice")
}

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, length float64
}

type rectangle2 struct {
	width, length float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

func (r2 rectangle2) area() float64 {
	return r2.width * r2.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.length)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("area = %f\n", g.area())
	fmt.Printf("perimeter = %f\n", g.perim())
}

func myPrinter(i ...interface{}) {
	for _, val := range i {
		fmt.Println(val)
	}
}

func printType(i interface{}) { // any is alias for interface{}
	switch i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: unknown")
	}
}
