package main

import "fmt"

func main() {

	// Multi dimensional array can be declared
	var numbers = [2][2]int{
		{1, 2}, {3, 4}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(numbers[i][j])
		}
	}

	//Accessing specific number in multi dimensional array
	fmt.Println(numbers[1][0])
}
