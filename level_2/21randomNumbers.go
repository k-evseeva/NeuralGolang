package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main_21() {

	fmt.Println(rand.Intn(101))     // from 0 to 100
	fmt.Println(rand.Intn(100) + 1) // from 1 to 100
	fmt.Println(rand.Intn(6) + 5)   // from 5 to 10

	val := rand.New(rand.NewSource(42))
	fmt.Println(val.Intn(101)) // 97 all the time

	val = rand.New(rand.NewSource(40))
	fmt.Println(val.Intn(101)) // 7 all the time

	val = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(val.Intn(101))

	fmt.Println(rand.Float64()) // from 0.0 to 1.0 (not including 1.0)

	// Dice game (2 dices)
	for {
		// Show menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter option (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice. Please, enter 1 or 2.")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing! See you soon!")
			break
		}

		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1

		// Show result
		fmt.Printf("You rolled %d and %d\n", dice1, dice2)
		fmt.Printf("Total: %d\n", dice1+dice2)

		// Ask if user wants to roll again
		fmt.Print("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)

		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no. Bye!")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing! See you soon!")
			break
		}
	}
}
