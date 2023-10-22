package main

import "fmt"

func main() {
	/* variable is declared with
	var variable_name variable_data_type
	variable name if started with capital letter means it's exposed
	if lower letter than it's not exposed outside the package
	*/
	var Name string = "Dinesh Ahuja"

	// data type if not provided , compiler will figure out during compile time
	var firstname = "Dinesh"

	// static declaration , means variable is declared in start and can be given a value during run time
	var lastname string
	lastname = "Ahuja"

	// multiple value declaration in a single line
	var age, gender = 26, "Male"

	// Dynamic declaration with := symbol
	address := "Xyz"

	var height float32 = 173.0

	var flag bool = true

	fmt.Printf("full name %v , type %T \n", Name, Name)
	fmt.Printf("first name %v , type %T \n", firstname, firstname)
	fmt.Printf("full name %v , type %T \n", lastname, lastname)
	fmt.Printf("Age %v , type %T \n", age, age)
	fmt.Printf("Gender %v , type %T \n", gender, gender)
	fmt.Printf("Address %v , type %T \n", address, address)
	fmt.Printf("Height %v , type %T \n", height, height)
	fmt.Printf("Flag %v , type %T \n", flag, flag)

}
