Go lang Constants


ou can use const prefix to declare constants with a specific type as follows −

const variable type = value;
The following example shows how to use the const keyword −


package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d", area)   
}