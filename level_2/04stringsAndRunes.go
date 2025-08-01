package main

import (
	"fmt"
	"unicode/utf8"
)

func main_04() {
	// creating strings
	message := "Hello\nGo!"
	message2 := "Hello\tGo!"
	message3 := "Hello\rGo!"
	rawMessage := `Hello\tGo!\n`
	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)
	// len of string
	fmt.Println("Length of message:", len(message))
	fmt.Println("Length of rawMessage:", len(rawMessage))
	// accessing by index
	fmt.Println("H =", message[0])
	fmt.Println("e =", message[1])
	fmt.Println("l =", message[2])
	fmt.Println("l =", message[3])
	fmt.Println("o =", message[4])
	// concatenation
	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting, name)
	fmt.Println(greeting + name)
	// comparison
	str1 := "apple" // a = 97 (ASCII)
	str2 := "Apple" // A = 65 (ASCII)
	str3 := "apple"
	str4 := "app"
	fmt.Println("str1 > str2:", str1 > str2)
	fmt.Println("str1 == str3:", str1 == str3)
	fmt.Println("str1 >= str3:", str1 >= str3)
	fmt.Println("str1 != str2:", str1 != str2)
	fmt.Println("str4 < str1:", str4 < str1)
	// iteration over string
	str := "Hello"
	for ind, char := range str {
		fmt.Printf("index: %d char: %c (code %d, 0x%x)\n", ind, char, char, char)
		fmt.Println("index:", ind, "char:", char)
	}
	// count runes
	fmt.Println("Length of str:", len(str))
	fmt.Println("Count runes:", utf8.RuneCountInString(str))
	// str[0] = "h" - error (strings are immutable)
	fullGreeting := greeting + name
	fmt.Println(fullGreeting)
	fullGreeting = greeting + ", " + name + "!"
	fmt.Println(fullGreeting)
	// runes
	var ch rune = 'a'
	ch2 := 'b'
	ch3 := 'ю'
	ch4 := 'я'
	fmt.Println(ch, ch2)
	fmt.Printf("%c %c\n", ch, ch2)
	fmt.Println(ch3, ch4)
	// cast to string
	chString := string(ch)
	fmt.Printf("chString = %s has type %T\n", chString, chString)
	fmt.Printf("ch = %c has type %T\n", ch, ch)
}
