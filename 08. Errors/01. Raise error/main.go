package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	value, err := Sqrt(-1)
	if err != nil {
		fmt.Println("Error : ", err)

	} else {
		fmt.Println("Sqrt value is ", value)
	}

	value, err = Sqrt(9)
	if err != nil {
		fmt.Println("Error : ", err)

	} else {
		fmt.Println("Sqrt value is ", value)
	}
}

func Sqrt(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("Number cannot be less than zero")
	} else {
		return math.Sqrt(float64(n)), nil
	}
}
