package main

import "fmt"

func main() {
	var a = 100
	var b = 200
	fmt.Println("a = ", a, " b = ", b)
	swapByValue(a, b)
	fmt.Println("a = ", a, " b = ", b)
	swapByReference(&a, &b)
	fmt.Println("a = ", a, " b = ", b)
}

// This is call by Value
func swapByValue(a int, b int) {
	var temp = b
	b = a
	a = temp
}

func swapByReference(a *int, b *int) {
	var temp int
	temp = *b
	*b = *a
	*a = temp
}
