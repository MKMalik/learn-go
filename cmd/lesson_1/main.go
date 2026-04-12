package main

import (
	"errors"
	"fmt"
)

func main() {
	var res int
	var remainder int
	var err error
	res, remainder, err = divide(5, 6)
	switch {
	case err != nil:
		fmt.Printf("Error occurred: %v\n", err)
	default:
		fmt.Printf("Result: %v\n", res)
		switch remainder {
		case 0:
			fmt.Println("Exact division")
		case 1, 2:
			fmt.Println("Close division")
		default:
			fmt.Println("Not close enough")
		}
	}
}

func divide(a int, b int) (int, int, error) {
	if b == 0 {
		var err error = errors.New("Can't divide by 0")
		return 0, 0, err
	}
	var division int = a / b
	var remainder int = a % b
	return division, remainder, nil
}
