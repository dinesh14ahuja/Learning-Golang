package main

import (
	"errors"
	"fmt"
)

// Create a custom Error struct
type CustomError struct {
	Message string
	status  int
}

// Custom error should override a function Error which should return string
func (c *CustomError) Error() string {
	return c.Message
}

// This function will create new object of custom error and return Custom error
func New(message string, status int) error {
	return &CustomError{message, status}
}

func main() {
	var a int = 5
	var b int = 10
	result, e := Subtract(a, b)
	if e != nil {
		// A null pointer to CustomError
		var customError *CustomError

		// errors.As function will convert the error to CustomError , to access all fields
		if errors.As(e, &customError) {
			fmt.Println(customError.Message, customError.status)
		}

	} else {
		fmt.Println(result)
	}
}

/*
Created a simple function to understand Custom error
If result of subtraction is negative then a custom error should be returned
*/
func Subtract(a int, b int) (int, error) {
	if a-b < 0 {
		return 0, New("Negative value", 400)
	} else {
		return (a - b), nil
	}
}
