package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, i)
	}
}
func main() {
	// go keyword will run the function in background and will create a goroutine
	// i.e a light weight thread
	go display("Dinesh")

	// This function works on main goroutine thread
	display("Gaurav")

	// time.sleep is applied to see it concurrently working
	// Otherwise if main thread is completed then the program will end

}
