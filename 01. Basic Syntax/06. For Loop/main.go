package main

import "fmt"

func main() {
	var a int
	var b int = 15

	numbers := [6]int{1, 2, 3, 4}

	for a = 0; a < 10; a++ {
		fmt.Println("value of a = ", a)
	}

	for a < b {
		a++
		fmt.Println("value of a = ", a)
	}

	for index, value := range numbers {
		fmt.Println("Value = ", value, ", at index = ", index)
	}
}
