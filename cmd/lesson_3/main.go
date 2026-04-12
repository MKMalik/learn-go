package main

import (
	"fmt"
)

func main() {
	var myMap map[string]uint8
	myMap = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"mkmalik": 26}
	myMap2["hello"] = 100
	fmt.Println(len(myMap2))
	fmt.Println(myMap2)
	for key, val := range myMap2 {
		fmt.Println(key, val)
	}
}
