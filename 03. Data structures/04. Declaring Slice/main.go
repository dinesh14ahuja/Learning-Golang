package main

import "fmt"

func main() {
	// A slice which is unsized
	var a []int

	fmt.Println("size of slice is ", len(a))
	fmt.Println("capacity of slice is ", cap(a))

	if a == nil {
		fmt.Println(" slice is nil ")
	}

	a = append(a, 1, 2, 3, 4, 5, 6, 7)

	PrintSlice(a)

	// A slice of specific len and capacity

	var b = make([]int, 5, 10)
	for i := 0; i < 5; i++ {
		b[i] = i * 2
	}

	fmt.Println("size of slice is ", len(b))
	fmt.Println("capacity of slice is ", cap(b))
	// The slice can be filled till capacity is reached , after the capacity it will resize it to double the capacity
	b = append(b, 10, 12, 14, 16, 18)
	fmt.Println("size of slice is ", len(b))
	fmt.Println("capacity of slice is ", cap(b))

	// Here we can see , on the insertion of 11th element , the capacity reached the limit of 10
	// Then resizing was done dynamically i.e old capacity + 10 which was defined during declaration
	b = append(b, 20)
	fmt.Println("size of slice is ", len(b))
	fmt.Println("capacity of slice is ", cap(b))
	PrintSlice(b)

}

func PrintSlice(a []int) {
	for index, value := range a {
		fmt.Println("index = ", index, " value = ", value)
	}
}
