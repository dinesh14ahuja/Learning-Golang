package main

import (
	"fmt"
	// packages-module is a mod
	// Referening from go.mod file
	// Created using go mod init command

	calculator "packages-module/Package"
)

func main() {
	var a = 20
	var b = 40

	c := calculator.Add(a, b)
	fmt.Println(c)

	// As we can see subtract function is not detected
	// Because to expose a function and a variable it should start with uppercase letter
	//d := calculator.subtract
}
