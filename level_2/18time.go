package main

import (
	"fmt"
	"time"
)

func main_18() {

	// Current local time
	fmt.Println(time.Now())

	// Specific time
	specTime := time.Date(2024, time.May, 30, 12, 55, 0, 1, time.UTC)
	fmt.Println(specTime)

	// Parse time
	parsedTime1, _ := time.Parse("2006-01-02", "2025-03-04")
	fmt.Println(parsedTime1)
	parsedTime2, _ := time.Parse("06-01-02", "25-03-04")
	fmt.Println(parsedTime2)
	parsedTime3, _ := time.Parse("06-1-2", "25-3-4")
	fmt.Println(parsedTime3)
	parsedTime4, _ := time.Parse("06-1-2 15-04", "25-3-4 18-03")
	fmt.Println(parsedTime4)

	// Formatting time
	tm := time.Now()
	fmt.Println("Formatted time:", tm.Format(" Monday 2006-01-02 15-04-05"))

	oneDayLater := tm.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded time:", tm.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	tm = time.Date(2025, time.May, 8, 14, 16, 40, 00, time.UTC)

	// Convert this to specific time zone
	tmLocal := tm.In(loc)

	// Perform rounding
	roundedTime := tm.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original time (UTC):", tm)
	fmt.Println("Original time (Local):", tmLocal)
	fmt.Println("Rounded time (UTC):", roundedTime)
	fmt.Println("Rounded time (Local):", roundedTimeLocal)
	fmt.Println("Truncated time (UTC):", tm.Truncate(time.Hour))

	// Convert time to location
	// cd $(go env GOROOT)/lib/time
	loc, _ = time.LoadLocation("America/New_York")
	tmInNY := time.Now().In(loc)
	fmt.Println("New York time:", tmInNY)

	// Find duration
	tm1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	tm2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := tm2.Sub(tm1)
	fmt.Println("Duration:", duration)

	// Compare time
	fmt.Println("tm2 is after tm1:", tm2.After(tm1))

}
