Go lang switch case

A switch statement allows a variable to be tested for equality against a list of values. Each value is called a case, and the variable being switched on is checked for each switch case.

In Go programming, switch statements are of two types −

Expression Switch − In expression switch, a case contains expressions, which is compared against the value of the switch expression.

Type Switch − In type switch, a case contain type which is compared against the type of a specially annotated switch expression.

Expression Switch
The syntax for expression switch statement in Go programming language is as follows −

switch(boolean-expression or integral type){
   case boolean-expression or integral type :
      statement(s);      
   case boolean-expression or integral type :
      statement(s); 
   
   /* you can have any number of case statements */
   default : /* Optional */
      statement(s);
}


Example

package main

import "fmt"

func main() {
   /* local variable definition */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }
   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n", grade );      
}


Type Switch
The syntax for a type switch statement in Go programming is as follows −

switch x.(type){
   case type:
      statement(s);      
   case type:
      statement(s); 
   /* you can have any number of case statements */
   default: /* Optional */
      statement(s);
}