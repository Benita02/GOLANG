package main

import "fmt"

func main() {
//ARRAYS FROM DAILY CODE BUFFER YOUTUBE CHANNEL
	//var array [3]int = [3]int{10, 20, 30}

	var arr []int = []int{1,2,3,4,5,6,7,8,9,10}
	//arr[0] = 36
	//SLICING OPERATIONS
	a := arr[:] //copies entire array 'arr' into variable a
	b := a[2:] //copies elements of a from the 2nd element to the end of the array
	c := b[:5] //Copies elements from start to the fifth element
	d := c[2:5] //Guess? from the 2nd to the 5th element
	e := arr[:len(arr)-5] 
	// dynamic := [...]int {586, 43, 44, 3, 4, 4, 55, 4, 67, 8, 6, 8, 5, 6}
	fmt.Printf("Array comprises of arr: %v\n", arr)
	fmt.Printf("Array comprises of A: %v\n", a)
	fmt.Printf("Array comprises of B: %v\n", b)
	fmt.Printf("Array comprises of C: %v\n", c)
	fmt.Printf("Array comprises of D: %v\n", d)
	fmt.Printf("Array comprises of E: %v\n", e)
	// fmt.Printf("Array comprises of: %v\n", dynamic)
	// fmt.Printf(" %v\n", len(array))
	// fmt.Printf(" %v\n", len(amt))
	// fmt.Printf(" %v\n", len(dynamic))
	
















	// fmt.Println("Benita is soooo cool")
	// var name string = "George of the Jungle"
	// fmt.Println(name)
	// name2 := "Rohit"
	// fmt.Println(name2)
	// name2 = "Da new name ain't Rohit but Benita"

	// var num int = 3
	// fmt.Println(num)

	// num = 76
	// fmt.Println(num)

	// var flag bool = true
	// fmt.Println(flag)

	// var pie float32 = 3.14
	// fmt.Println(pie)

	// var age int = -89
	// // if age > 0 && age <= 2 {
	// // 	fmt.Println("You are an infant")
	// // } else if age > 2 && age <= 5 {
	// // 	fmt.Println("You are a toddler")
	// // } else if age > 5 && age <= 12 {
	// // 	fmt.Println("You are a child")
	// // } else if age > 12 && age <= 18 {
	// // 	fmt.Println("You are a teenager, you stubborn brat")
	// // } else if age > 18 && age <= 35 {
	// // 	fmt.Println("You're a bloody adult, go get a job")
	// // } else if age > 35 && age <= 49 {
	// // 	fmt.Println("You're middle aged, hope you made the right choices")
	// // } else if age > 50 {
	// // 	fmt.Println("You're an elder, congrats, this wicked world didn't kill you......YET!")
	// // } else {
	// // 	fmt.Println("Invalid Input, try your shit again, dumbass!")
	// // }

	// switch {
	// case age > 0 && age <= 2:
	// 	fmt.Println("You are an infant")
	// case age > 2 && age <= 5:
	// 	fmt.Println("You are a toddler")
	// case age > 5 && age <= 12:
	// 	fmt.Println("You are a child")
	// case age > 12 && age <= 18:
	// 	fmt.Println("You are a teenager, you stubborn brat")
	// case age > 18 && age <= 35:
	// 	fmt.Println("You're a bloody adult, go get a job")
	// case age > 35 && age <= 49:
	// 	fmt.Println("You're middle aged, hope you made the right choices")
	// case age > 50:
	// 	fmt.Println("You're an elder, congrats, this wicked world didn't kill you......YET!")
	// default:
	// 	fmt.Println("Invalid Input, try your shit again, dumbass!")

	// }
}
