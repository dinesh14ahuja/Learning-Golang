package main

import "fmt"

func main() {
	var a int = 42

	// Declaring a pointer in golang has to be for a specific type
	// *int means a pointer for type int
	// So this pointer can only point to a int type variable
	var p *int

	// Pointer contains memory address of a variable of similar type
	// address of a variable is obtained using & operator
	p = &a

	fmt.Println("address of a ", &a)
	fmt.Println("pointer value of p ", p)

	// Since P contains address of a
	// We can dereference the pointer to obtain actual value on that address
	// dereference is done using * operator on the variable
	fmt.Println("Value dereferenced by p ", *p)

	// Similarly we can change the value on that address
	*p = 24

	fmt.Println("Value of a ", a)

}
