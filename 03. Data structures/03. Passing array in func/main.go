package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	average := getAverage(a, len(a))
	fmt.Printf("Average value of %v = %v", a, average)
}

/*

While defining a function , if parameter is [] int
means Unsized array . Then the declaration should also be unsized.

If we define a function with [5]int then the declaration should also be

var a = [5]int{1,2,3,4,5}
*/
func getAverage(array []int, size int) int {

	var sum, average int
	for _, value := range array {
		sum += value
	}
	average = sum / size
	return average
}
