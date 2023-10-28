package main

import "fmt"

// A new type is defined of name Person which is a struct of below fields
type Person struct {
	Name    string
	Age     int
	Gender  string
	address string
}

func main() {

	var person1 Person
	person1 = Person{"Dinesh", 26, "M", "XYZ "}
	// %+v is used to print the struct with field and value
	fmt.Printf("%+v\n", person1)

	// We can access individual fields of struct
	fmt.Println(person1.Age)
	fmt.Println(person1.Name)
	fmt.Println(person1.Gender)
	fmt.Println(person1.address)

	// If some fields are not passed then they will be initialized to default value
	person2 := Person{Name: "Gaurav", Age: 28}
	fmt.Printf("%+v\n", person2)
}
