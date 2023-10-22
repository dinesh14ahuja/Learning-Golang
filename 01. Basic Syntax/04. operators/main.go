package main

import "fmt"

func main() {
	//Arithmetic Operators

	var a = 10
	var b = 20

	fmt.Printf("a+b=%v\n", (a + b))
	fmt.Printf("a-b=%v\n", (a - b))
	fmt.Printf("a*b=%v\n", (a * b))
	fmt.Printf("a/b=%v\n", (a / b))
	fmt.Printf("a/b remainder=%v\n", (a % b))

	//Relational Operator

	var c = 20
	var d = 10
	var e = 20

	fmt.Printf("c==d is %v\n", (c == d))
	fmt.Printf("c!=d is %v\n", (c != d))
	fmt.Printf("c>=d is %v\n", (c >= d))
	fmt.Printf("c<=d is %v\n", (c <= d))
	fmt.Printf("c==e is %v\n", (c == e))
	fmt.Printf("c>e is %v\n", (c > e))
	fmt.Printf("c<e is %v\n", (c < e))
	fmt.Printf("c!=e is %v\n", (c != e))

	//Logical Operators

	var g = true
	var h = false

	fmt.Printf("g&&h = %v\n", (g && h))
	fmt.Printf("g||h = %v\n", (g || h))
	fmt.Printf("!(g&&h) = %v\n", !(g && h))

}
