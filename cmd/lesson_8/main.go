package main

import "fmt"

func main() {
	var sliceInt []int = []int{3,1,23,40,12}
	var sliceFloat32 []float32 = []float32{4.9, 5.0, 21.999999}
	fmt.Println(sumSlice(sliceInt))
	fmt.Println(sumSlice(sliceFloat32))
	fmt.Println(isEmpty(sliceFloat32))
}

func sumSlice[T int | float32 | float64] (slice[] T) T {
	var sum T
	for _, num := range slice {
		sum += num
	}
	return sum
}

func isEmpty[T any] (slice[] T) bool {
	return len(slice) == 0
}
