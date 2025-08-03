package main

import (
	"errors"
	"fmt"
)

func main_13() {
	custErr := doSomething()
	if custErr != nil {
		fmt.Println(custErr)
		return
	}
	fmt.Println("Operation completed successfully")
}

type customError struct {
	code    int
	message string
	er      error
}

func (c *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", c.code, c.message, c.er)
}

// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong",
// 	}
// }

func doSomethingElse() error {
	return errors.New("internal error")
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}
