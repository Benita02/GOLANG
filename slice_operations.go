package main

import "fmt"

func main() {
	var a []int = []int{1, 2 ,3} //when you don't provide a size, your using a slice operator.
	var b []int = a // This is actually pointing var b to the memory location of a and not the elements,
	//so if you change b, you change the mem. location of a, meaning that changing b affects a
	fmt.Println(a)
	fmt.Println(b)
	b[0] = 73
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(a))
	fmt.Println(cap(b))// Cap function to get capacity of the slice

	//USING MAKE FUNCTION TO CREATE SLICES
	var aa []int = make([]int, 3, 10)//make function, argument 3 - length and 10 - capacity
	aa = append(aa, 36, 800, 43)
	fmt.Println(aa)

	//USING APPEND() FUNCTION (takes in 2 arguments, slice you want to append to, and what you want to append)
	var bb []int = append(aa, 99)
	fmt.Println(bb)

	var cc []int = append(bb, aa...)
	fmt.Println(cc)

}