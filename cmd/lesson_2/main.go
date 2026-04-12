package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [3]int32 = myArrayFunc()
	fmt.Println(&array[0], &array[1], &array[2])
}

func myArrayFunc() [3]int32 {
	var myArray [3]int32
		var i int = 0
	for i < 3 {
		myArray[i] = int32(rand.Int())
		i++
	}
	return myArray
}
