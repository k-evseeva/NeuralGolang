package main

import (
	"fmt"
	"os"
)

func main_23() {
	defer fmt.Println("Deferred statement")

	// only this will be executed
	fmt.Println("Start of main function")
	os.Exit(1)

	fmt.Println("End of main function")
}
