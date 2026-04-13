package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rectangle := Rectangle{height: 5, width: 4}
	fmt.Println(calculateArea(rectangle))

	circle := Circle{radius: 2}
	fmt.Println(calculateArea(circle))
}
