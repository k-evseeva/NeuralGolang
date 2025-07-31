package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main_09() {
	// guessing game
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")
	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Less... Try guessing a greater number!")
		} else {
			fmt.Println("Greater... Try guessing a less number!")
		}
	}
}
