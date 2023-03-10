// In Go, the defer keyword is used to schedule a function call to be executed just before the function that contains the defer statement returns.
// The function call is deferred until the surrounding function completes, regardless of whether it completes normally or with an error.
// This is useful in situations where you want to ensure that certain operations are performed before a function exits, 
// regardless of what happens inside the function. For example, you may want to close a file or release a lock before
// returning from a function, to avoid leaking resources.The defer keyword can also be used to simplify code 
// by postponing certain cleanup tasks until the end of a function. 
// It is important to note that deferred functions are executed in a
// last-in, first-out (LIFO) order, meaning that
// the last deferred function will be executed first,
// followed by the second-to-last, and so on.


package main 

import "fmt"

func main(){
	// fmt.Println(1)
	// defer fmt.Println(2) //defer statement will execute after every other code in it's block runs
	// fmt.Println(3)
//if they were all defer statements, they would print out in reverse i.e 3,2,1 cause whatever's added last comes out first

//Whenever there's a panic created the rest of the code won't execute
	fmt.Println("START")
	panic("This is  a panic")
	fmt.Println("END")//Line does'nt run 

	defer fmt.Println(8)
	fmt.Println(2) //defer statement will execute after every other code in it's block runs
	fmt.Println(3)
}