package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func main() {
	person1 := Person{"Dinesh", 26, "M", "XYZ"}
	person2 := Person{"Gaurav", 28, "M", "ABC"}
	person3 := Person{"Gauri", 26, "F", "DEF"}

	checkGender(person1)
	checkGender(person2)
	checkGender(person3)

	checkWhoIsOlder(person1, person2)
	checkWhoIsOlder(person1, person3)
	checkWhoIsOlder(person2, person3)
}

func checkGender(person Person) {
	if person.Gender == "M" {
		fmt.Printf("%v is a Male\n", person.Name)
	} else {
		fmt.Printf("%v is a Female\n", person.Name)
	}
}

func checkWhoIsOlder(person1 Person, person2 Person) {
	if person1.Age > person2.Age {
		fmt.Printf("%v is older than %v\n", person1.Name, person2.Name)
	} else if person1.Age < person2.Age {
		fmt.Printf("%v is older than %v\n", person2.Name, person1.Name)
	} else {
		fmt.Printf("%v Age is same as %v\n", person1.Name, person2.Name)
	}

}
