package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	additionaResult := add(a, b)
	fmt.Println("Addition = ", additionaResult)

	subtractionResult := subtract(a, b)
	fmt.Println("Subration = ", subtractionResult)

	division, remainder := division(a, b)
	fmt.Println("division = ", division, ", remainder=", remainder)
}

// function has function name , parameters and return type
func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// Function can have multiple returns as well
// Mostly used to return error
func division(a int, b int) (int, int) {
	return a / b, a % b
}
