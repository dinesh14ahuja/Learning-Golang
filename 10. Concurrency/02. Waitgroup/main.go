package main

import (
	"fmt"
	"sync"
	"time"
)

// Paramter should have pointer to waitgroup
func PrintName(s string, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(100)
	}
	// Once process is done then we should mark waitgroup as done
	wg.Done()
}

func PrintNumbers(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := 0; j < i; j++ {
		fmt.Println(j)
		time.Sleep(100)
	}
}

func main() {

	// Waitgroup is like an manager which manage the go routines

	var wg sync.WaitGroup
	// We should add how many go routines will run in parallel
	wg.Add(2)

	// Go routines should have a parameter to pointer of Waitgroup
	go PrintName("Dinesh", &wg)
	go PrintName("Gaurav", &wg)

	//wg.Add(1)
	// If we don't add waitgroup
	// then the goroutines will work concurrently but while calling wg.wait it will check for 0 count only
	go PrintNumbers(100, &wg)
	// In the end we can call waitgroup.wait . This function will wait until all go routines are done
	wg.Wait()
	fmt.Println("Main thread")
}
