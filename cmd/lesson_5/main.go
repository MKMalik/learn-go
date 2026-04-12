package main

import "fmt"

func main() {
	// var p *string = new(string)
	// var s string = *p
	// fmt.Println(p)
	// fmt.Println(&s)
	// p = &s
	// fmt.Println(p)
	// fmt.Println(&s)
	// s = "hello"
	// *p = "hello1"
	// fmt.Println(s, *p)
	var slices [3]int = [3]int{2, 3, 9}
	squareByRef(&slices)
	fmt.Printf("slices mem loc: %p\n", &slices)
}

func squareByRef(list *[3]int) {
	fmt.Printf("list mem loc:   %p\n", list)
	for i := range list {
		list[i] = list[i] * list[i]
	}
}
