Go programming language allows multidimensional arrays. 
Here is the general form of a multidimensional array declaration −

var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
For example, the following declaration creates a three dimensional 5 . 10 . 4 integer array −

var threedim [5][10][4]int



Two-Dimensional Arrays
A two-dimensional array is the simplest form of a multidimensional array. A two-dimensional array is, in essence, a list of one-dimensional arrays. To declare a two-dimensional integer array of size [x, y], you would write something as follows −

var arrayName [ x ][ y ] variable_type
Where variable_type can be any valid Go data type and arrayName will be a valid Go identifier. A two-dimensional array can be think as a table which will have x number of rows and y number of columns. A 2-dimensional array a, which contains three rows and four columns can be shown as below −

Two Dimensional Arrays in Go
Thus, every element in the array a is identified by an element name of the form a[ i ][ j ], where a is the name of the array, and i and j are the subscripts that uniquely identify each element in a.

Initializing Two-Dimensional Arrays
Multidimensional arrays may be initialized by specifying bracketed values for each row. Following is an array with 3 rows and each row has 4 columns.

a = [3][4]int{  
   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   {8, 9, 10, 11}   /*  initializers for row indexed by 2 */
}
Accessing Two-Dimensional Array Elements
An element in two dimensional array is accessed by using the subscripts, i.e., row index and column index of the array. For example −

int val = a[2][3]