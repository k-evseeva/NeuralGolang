package main

import (
	"fmt"
	"time"
)

func main_20() {

	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00"
	val := "2025-08-06T14:30:18Z"

	t, err := time.Parse(layout, val)

	if err != nil {
		fmt.Println("Error parsinf time:", err)
		return
	}

	fmt.Println(t)

	layout2 := "Jan 02, 2006 03:04 PM"
	val2 := "Jul 10, 2025 03:18 PM"

	t2, _ := time.Parse(layout2, val2)

	fmt.Println(t2)

}
