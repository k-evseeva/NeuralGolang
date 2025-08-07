package main

import (
	"fmt"
	"time"
)

func main_19() {

	// 00:00:00 UTC on Jan 1, 1970
	// year 2038 problem
	now := time.Now()
	unix := now.Unix()
	fmt.Println("now:", now)
	fmt.Println("unix (timestamp):", unix)
	utm := time.Unix(unix, 0)
	fmt.Println("unix (time):", utm)
	fmt.Println("Time:", utm.Format("2006-01-02"))
	fmt.Println(time.Since(now))

}
