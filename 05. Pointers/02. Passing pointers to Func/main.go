package main

import "fmt"

func main() {

	// A very basic example of how to pass pointers in parameters
	var sum int = 0
	var a int = 20
	var b int = 30

	// Instead of passing variable values , we pass there address
	// This is refered as Pass by Reference
	Sum(&a, &b, &sum)

	// here we can see variable sum is updated by the function
	fmt.Println(sum)

	fmt.Println("before swap value of a ", a)
	fmt.Println("before swap value of b ", b)

	Swap(&a, &b)

	fmt.Println("before swap value of a ", a)
	fmt.Println("before swap value of b ", b)

}

func Swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

//The function takes pointers as parameters
func Sum(a *int, b *int, sum *int) {
	// Here pointers are first derefered first and then sum pointer is updated
	*sum = *a + *b
}
