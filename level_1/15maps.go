package main

import (
	"fmt"
	"maps"
)

func main_15() {
	// creating map with make
	myMap := make(map[string]int)
	fmt.Println(myMap)
	// accessing by key
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["unknownKey"])
	myMap["code"] = 21
	fmt.Println(myMap)
	// delete element
	delete(myMap, "code")
	fmt.Println(myMap)
	// add new elements
	myMap["key2"] = 2
	myMap["key3"] = 3
	myMap["key4"] = 4
	myMap["key5"] = 5
	fmt.Println(myMap)
	// delete all elements
	clear(myMap)
	fmt.Println(myMap)
	// checking for existence
	val, ok := myMap["key1"]
	fmt.Println(val, ok)
	myMap["key1"] = 1
	val, ok = myMap["key1"]
	fmt.Println(val, ok)
	_, ok = myMap["key1"]
	fmt.Println("Is any value associated with key1:", ok)
	// creating new maps
	myMap2 := map[string]int{"a": 1, "b": 2}
	myMap3 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	// check for equal maps
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("myMap and myMap2 are not equal")
	}
	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	} else {
		fmt.Println("myMap2 and myMap3 are not equal")
	}
	// iteration over map
	for key, val := range myMap3 {
		fmt.Println(key, val)
	}
	for _, val := range myMap3 {
		fmt.Println(val)
	}
	// create nil map
	var myMap4 map[string]string
	fmt.Println("Is myMap4 is initialized with nil:", myMap4 == nil)
	val4 := myMap4["key1"]
	fmt.Printf("|%s|\n", val4)
	// myMap4["key1"] = "val1" - error
	fmt.Println(myMap4)
	myMap4 = make(map[string]string)
	myMap4["key1"] = "val1" // no error
	fmt.Println(myMap4)
	fmt.Println(len(myMap4))
	// nested map
	nestedMap := make(map[string]map[string]int)
	nestedMap["map1"] = myMap
	nestedMap["map2"] = myMap2
	fmt.Println(nestedMap)
	fmt.Printf("%v\n", myMap4)
}
