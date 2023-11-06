package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateFile(filename string, data string) {
	// os library is used to do Operating system related operations
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error occured while creating file")
		log.Fatalf("Error : %s", err)
	}

	// defer will close the file in the end
	defer file.Close()

	len, err := file.WriteString(data)
	if err != nil {
		fmt.Println("Error occured while writing a file")
		log.Fatalf("Error : %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)

}

func Readfile(filename string) {

	// The os package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}
func main() {
	fmt.Println("Enter Filename")
	var filename string
	fmt.Scanln(&filename)

	fmt.Println("Enter text to add...")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error : %s", err)
	}
	CreateFile(filename, input)
	Readfile(filename)

}
