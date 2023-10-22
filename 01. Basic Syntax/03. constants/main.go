package main

import "fmt"

func main() {
	//const are variables whose values will be fixed and cannot be changed.
	const name string = "Dinesh"

	//static binding is not allowed. const has to be declared and initiated in same line.
	// dynamic binding is also not allowed

	const length = 20
	const width = 40

	var area int = length * width

	fmt.Printf("Area is %d\n", area)

	// Mutiple const can be declared in a ()
	// iota is a unsigned integer which will increment in a const block
	const (
		a = iota
		b = 2 * iota
		c = "Dinesh"
	)
	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)
	fmt.Printf("%v , %T\n", c, c)
	// Iota cannot be used for a variable
	//var d = iota
	// Iota reset to 0 for new const block
	const (
		d = iota
		e = 2 * iota
		f = "Dinesh"
	)
	fmt.Printf("%v , %T\n", d, d)
	fmt.Printf("%v , %T\n", e, e)
	fmt.Printf("%v , %T\n", f, f)
}
