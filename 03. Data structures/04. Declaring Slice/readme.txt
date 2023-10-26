Go Slice is an abstraction over Go Array. Go Array allows you to define variables that can hold several data items of the same kind but it does not provide any inbuilt method to increase its size dynamically or get a sub-array of its own. Slices overcome this limitation. It provides many utility functions required on Array and is widely used in Go programming.

Defining a slice
To define a slice, you can declare it as an array without specifying its size. Alternatively, you can use make function to create a slice.

var numbers []int /* a slice of unspecified size */
/* numbers == []int{0,0,0,0,0}*/
numbers = make([]int,5,5) /* a slice of length 5 and capacity 5*/

len() and cap() functions
A slice is an abstraction over array. It actually uses arrays as an underlying structure. The len() function returns the elements presents in the slice where cap() function returns the capacity of the slice (i.e., how many elements it can be accommodate). The following example explains the usage of slice −


package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)
   printSlice(numbers)
}
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
When the above code is compiled and executed, it produces the following result −

len = 3 cap = 5 slice = [0 0 0]



Nil slice
If a slice is declared with no inputs, then by default, it is initialized as nil. Its length and capacity are zero