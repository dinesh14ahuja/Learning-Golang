If you want to pass a single-dimension array as an argument in a function, you would have to declare function formal parameter in one of following two ways and all two declaration methods produce similar results because each tells the compiler that an integer array is going to be received. Similar way you can pass multi-dimensional array as formal parameters.

Way-1
Formal parameters as a sized array as follows −

void myFunction(param [10]int)
{
.
.
.
}
Way-2
Formal parameters as an unsized array as follows −

void myFunction(param []int)
{
.
.
.
}