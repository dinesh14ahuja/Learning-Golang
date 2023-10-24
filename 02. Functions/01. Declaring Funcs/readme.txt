Declaring Funcs in Go Lang

A function is a group of statements that together perform a task. Every Go program has at least one function, which is main(). You can divide your code into separate functions. How you divide your code among different functions is up to you, but logically, the division should be such that each function performs a specific task.

A function declaration tells the compiler about a function name, return type, and parameters.
A function definition provides the actual body of the function.

func function_name( [parameter list] ) [return_types]
{
   body of the function
}


Func − It starts the declaration of a function.

Function Name − It is the actual name of the function. The function name and the parameter list together constitute the function signature.

Parameters − A parameter is like a placeholder. When a function is invoked, you pass a value to the parameter. This value is referred to as actual parameter or argument. The parameter list refers to the type, order, and number of the parameters of a function. Parameters are optional; that is, a function may contain no parameters.

Return Type − A function may return a list of values. The return_types is the list of data types of the values the function returns. Some functions perform the desired operations without returning a value. In this case, the return_type is the not required.

Function Body − It contains a collection of statements that define what the function does.
