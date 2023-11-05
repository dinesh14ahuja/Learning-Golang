package main

import "fmt"

func main() {
	mychan := make(chan string)
	anotherchan := make(chan string)

	go func() {
		mychan <- "Dog"
	}()

	go func() {
		anotherchan <- "Cat"
	}()

	// select is just like switch statement , where it will wait for anyone channel to receive data from.
	// Whichever data is received first
	select {
	case mychandata := <-mychan:
		fmt.Println(mychandata)
	case myanotherchandata := <-anotherchan:
		fmt.Println(myanotherchandata)
	}
}
