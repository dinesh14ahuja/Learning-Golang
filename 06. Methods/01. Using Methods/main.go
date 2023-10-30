package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

// Method in Golang is terminology for using context
// Here this function specifies that the function Area can only work in context of struct Circle

// circle variable is called Receiver

// Methods work same as functions
// The main aim of methods is to implement interfaces
func (circle *Circle) Area() float64 {
	circle.radius += 1
	fmt.Printf("circle address->%p\n", circle)
	return math.Pi * circle.radius * circle.radius
}

func AreaCircle(circle Circle) float64 {
	circle.radius -= 2
	fmt.Printf("circle address->%p\n", &circle)
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}

	fmt.Printf("circle address main->%p\n", &circle)
	fmt.Printf("Circle area: %f\n", circle.Area())
	fmt.Println("radius ", circle.radius)
	fmt.Printf("Circle area: %f\n", AreaCircle(circle))
	fmt.Println("radius ", circle.radius)

}
