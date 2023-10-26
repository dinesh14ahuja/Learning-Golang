package main

import "fmt"

func main() {
	var fruits = [5]string{"apple", "banana", "kiwi", "orange", "peach"}

	// Iterate over the array of fruits using for loop
	for _, value := range fruits {
		fmt.Println(value)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// We can access specific array index
	fmt.Println(fruits[1])

}
