package main

import "fmt"

func main() {

	fmt.Println(addNumbers(1, 2, 3, 4, 5, 6, 10))
}

func addNumbers(nums ...int) int {
	var sum int
	for _, value := range nums {
		sum += value
	}
	return sum
}
