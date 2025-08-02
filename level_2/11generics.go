package main

import "fmt"

func main_11() {
	x, y := 1, 3
	fmt.Printf("x=%v, y=%v\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("x=%v, y=%v\n", x, y)
	a, b := "John", "Davey"
	fmt.Printf("a=%v, b=%v\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("a=%v, b=%v\n", a, b)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(8)
	intStack.push(11)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Printf("Stack is empty: %t\n", intStack.isEmpty())
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Printf("Stack is empty: %t\n", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Printf("Stack is empty: %t\n", stringStack.isEmpty())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack elements: ")
	for _, val := range s.elements {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
