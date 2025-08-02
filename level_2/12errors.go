package main

import (
	"errors"
	"fmt"
	"math"
)

func main_12() {
	res, err := sqrt(-16)
	handling_error(res, err)
	res2, err2 := sqrt(16)
	handling_error(res2, err2)
	data := []byte{}
	err3 := process(data)
	handling_error(data, err3)
	data2 := []byte{1, 2, 3}
	err4 := process(data2)
	handling_error(data2, err4)
	err5 := eprocess()
	if err5 != nil {
		fmt.Println(err5)
	}
	err6 := readData()
	if err != nil {
		fmt.Println(err6)
	} else {
		fmt.Println("Data read successfully")
	}
}

func handling_error(data any, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result=%v\n", data)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Square root of negative number")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// process data
	fmt.Println("Data processed successfully")
	return nil
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %v", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readConfig() error {
	return errors.New("Config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err) // %w - for errors
	}
	return nil
}
