package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func process(c chan int) {
	defer fmt.Println("process func exit")
	defer close(c)
	for i := range 5 {
		c <- i + 1
	}
}

// deadlock
// func main() {
// 	var c = make(chan int)
// 	c <- 1
// 	var i = <- c
// 	fmt.Println(i)
// }
