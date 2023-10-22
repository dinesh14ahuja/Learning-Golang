package main

import "fmt"

type person struct {
	name string
}

func main() {
	/*
		Different data types in go lang
	*/

	var a int = 12
	var b uint = 10
	var c = 20.0
	var d float64 = 132.47827

	var e string = "xyz123"

	var f *int = &a

	var g bool = true

	var newperson person = person{"Dinesh"}

	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)
	fmt.Printf("%v , %T\n", c, c)
	fmt.Printf("%v , %T\n", d, d)
	fmt.Printf("%v , %T\n", e, e)
	fmt.Printf("%v , %T\n", f, f)
	fmt.Printf("%v , %T\n", *f, *f)
	fmt.Printf("%v , %T\n", g, g)
	fmt.Printf("%v , %T\n", newperson, newperson)
}
