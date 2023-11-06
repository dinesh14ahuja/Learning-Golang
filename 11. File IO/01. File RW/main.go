package main

import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {
	// os library is used to do Operating system related operations
	file, err := os.Create("testfile.txt")

	if err != nil {
		fmt.Println("Error occured while creating file")
		log.Fatalf("Error : %s", err)
	}

	// defer will close the file in the end
	defer file.Close()

	len, err := file.WriteString("A sample text file created by Go ")
	if err != nil {
		fmt.Println("Error occured while writing a file")
		log.Fatalf("Error : %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)

}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "test.txt"

	// The os package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := os.ReadFile("testfile.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func main() {
	CreateFile()
	ReadFile()
}
