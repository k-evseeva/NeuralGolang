package main

import (
	"fmt"
	"regexp"
)

func main_17() {

	fmt.Println("He said, `I am great`")
	fmt.Println("He said, \"I am great\"")

	// Compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test emails
	email1 := "user@email.com"
	email2 := "invalid_email@mail.r"

	// Match
	fmt.Println("Email1: ", re.MatchString(email1))
	fmt.Println("Email2: ", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test date
	date1 := "2025-07-18"

	// Find all submatches
	submatches := re.FindStringSubmatch(date1)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Source string
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - multiline model
	// s - dot matches all

	// Another test string
	str = "Golang is great"
	re = regexp.MustCompile(`(?i)go`)
	fmt.Println("Match:", re.MatchString(str))

}
