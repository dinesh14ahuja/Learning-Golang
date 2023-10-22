package main

import "fmt"

func main() {
	var a int = 10
	var b int = 5

	if a > b {
		fmt.Println("a is greater than b")
	}

	fmt.Printf("Value of a is %v\n", a)

	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	var c = 20

	if a > b {
		fmt.Println("a is greater than b")
		if a > c {
			fmt.Println("a is greater than b")
		} else {
			fmt.Println("c is greater than a")
		}
	}

	if a < b {
		fmt.Println("b is greater than a")
	} else if a < c {
		fmt.Println("c is greater than a")
	} else {
		fmt.Println("a is smallest")
	}
}
