Go Lang For loop 

There may be a situation, when you need to execute a block of code several number of times. In general, statements are executed sequentially: The first statement in a function is executed first, followed by the second, and so on.

Programming languages provide various control structures that allow for more complicated execution paths.

A loop statement allows us to execute a statement or group of statements multiple times and following is the general form of a loop statement in most of the programming languages −


A for loop is a repetition control structure. It allows you to write a loop that needs to execute a specific number of times.

Syntax
The syntax of for loop in Go programming language is −

for [condition |  ( init; condition; increment ) | Range] {
   statement(s);
}

The flow of control in a for loop is a follows −

If a condition is available, then for loop executes as long as condition is true.

If a for clause that is ( init; condition; increment ) is present then −

The init step is executed first, and only once. This step allows you to declare and initialize any loop control variables. You are not required to put a statement here, as long as a semicolon appears.

Next, the condition is evaluated. If it is true, the body of the loop is executed. If it is false, the body of the loop does not execute and the flow of control jumps to the next statement just after the for loop.

After the body of the for loop executes, the flow of control jumps back up to the increment statement. This statement allows you to update any loop control variables. This statement can be left blank, as long as a semicolon appears after the condition.

The condition is now evaluated again. If it is true, the loop executes and the process repeats itself (body of loop, then increment step, and then again the condition). After the condition becomes false, the for loop terminates.

If range is available, then the for loop executes for each item in the range.

