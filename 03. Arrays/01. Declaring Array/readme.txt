Go programming language provides a data structure called the array, which can store a fixed-size sequential collection of elements of the same type. 
An array is used to store a collection of data, but it is often more useful to think of an array as a collection of variables of the same type.

Instead of declaring individual variables, such as number0, number1, ..., and number99, 
you declare one array variable such as numbers and use numbers[0], numbers[1], and ..., numbers[99] to represent individual variables. A specific element in an array is accessed by an index.

All arrays consist of contiguous memory locations. The lowest address corresponds to the first element and the highest address to the last element.

Declaring Arrays
To declare an array in Go, a programmer specifies the type of the elements and the number of elements required by an array as follows −

var variable_name [SIZE] variable_type

