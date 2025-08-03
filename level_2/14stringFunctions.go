package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main_14() {

	// Basics
	str := "Hello, Go!"
	fmt.Println(len(str))
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
	fmt.Println(str[0])
	fmt.Println(str[1:8])

	// String convertion
	num := 18
	num_str := strconv.Itoa(num)
	fmt.Println(len(num_str))

	// String splitting
	fruits1 := "apple, orange, banana"
	parts1 := strings.Split(fruits1, ", ")
	fmt.Println(parts1)
	fruits2 := "apple-orange-banana"
	parts2 := strings.Split(fruits2, "-")
	fmt.Println(parts2)

	// String joining
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, "*")
	fmt.Println(joined)

	// Search for substring
	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Contains(str, "Go?"))

	// Replace substrings
	gogo := "Go, Go!"
	replaced1 := strings.Replace(gogo, "Go", "World", 1)
	replaced2 := strings.Replace(gogo, "Go", "World", 2)
	fmt.Println(replaced1)
	fmt.Println(replaced2)

	// Remove spaces at the beginning and at the end
	strWithSpaces := "   Hello Everyone! "
	fmt.Println(strWithSpaces)
	fmt.Println(strings.TrimSpace(strWithSpaces))

	// Change case
	fmt.Println(strings.ToLower(strWithSpaces))
	fmt.Println(strings.ToUpper(strWithSpaces))

	// Repeating string
	fmt.Println(strings.Repeat("bar", 3))

	// Counting substring
	fmt.Println(strings.Count("Hello", "l"))

	// Check prefix anf suffix of string
	fmt.Println(strings.HasPrefix("Hello", "H"))
	fmt.Println(strings.HasPrefix("Hello", "h"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "lo!"))

	// Some regulars
	str3 := "Hel1o, 123 Go! It's 456."
	re := regexp.MustCompile(`\d+`)       // has > 0 digits
	matches := re.FindAllString(str3, -1) // find all matches
	fmt.Println(matches)

	// Rune counter
	str4 := "Hello, 世界"
	fmt.Println(len(str4))
	fmt.Println(utf8.RuneCountInString(str4))

	// String builder (for long strings)
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")
	builder_res := builder.String()
	fmt.Println(builder_res)
	builder.WriteRune('!')
	builder.WriteRune(' ')
	builder.WriteString("It's time to wake up")
	builder.WriteRune('!')
	builder_res = builder.String()
	fmt.Println(builder_res)

	// Reset builder
	builder.Reset()
	builder.WriteString("Refreshed")
	builder_res = builder.String()
	fmt.Println(builder_res)
}
