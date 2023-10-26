package main

import "fmt"

func main() {
	// map is declared as syntax
	// map[Key]Value
	var nameAge = map[string]int{"Dinesh": 26, "Gaurav": 28}

	for k, v := range nameAge {
		fmt.Printf("Key %v,Value %v\n", k, v)
	}

	nameAge["Akash"] = 28
	for k, v := range nameAge {
		fmt.Printf("Key %v,Value %v\n", k, v)
	}

	// We can create map using make function
	var newMap = make(map[string]int)
	fmt.Println("length of newmap = ", len(newMap))
	for k, v := range newMap {
		fmt.Printf("Key %v,Value %v\n", k, v)
	}

	fmt.Println("Age of Dinesh is =", nameAge["Dinesh"])

	delete(nameAge, "Gaurav")
	for k, v := range newMap {
		fmt.Printf("Key %v,Value %v\n", k, v)
	}

	value, ok := nameAge["Anil"]
	if ok {
		fmt.Println("Age of Anil is ", value)
	} else {
		fmt.Println("Anil is not present in the map")
	}

}
