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
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.Area())

}
