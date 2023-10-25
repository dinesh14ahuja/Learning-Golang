package main

import (
	"fmt"
	"math"
)

func main() {

	//Anonymous function assigned to a varaiable
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(36))
}
