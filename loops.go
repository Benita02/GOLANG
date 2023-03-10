//Golang has only for loops
package main

 import "fmt"
 
 func main(){
	i, j := 0, 0
	for i < 5{
		fmt.Println("Value of i is less than 10")
		fmt.Println("\n")
		fmt.Println(i, j)
		i += 1
		j += 1
	}

//The 'range' keyword is used to iterate over elements in a variety of
// data structures such as arrays, slices, maps, and strings.
	a := []int{1,2,3,4,5,6}
	for k, v := range a { //remember we can use the underscore i.e _, v or k,_
		fmt.Println(k, v)
	}

//Looping through a map
	friends_colleagues := make(map[int][]string, 5 )
	friends_colleagues = map[int][]string{
		100 : {"Royce"},
		200 : {"Siri,", "Maro,", "Lucky"},
		300 : {"Daniella,", "Laju,", "David,", "Tegha,", "MJ"},
		400 : {"Ephraim,", "Elizabeth"},
		500 : {"Cynthia,","Yahaya"},
		}
		for k := range friends_colleagues{
			fmt.Println(k)
		}
 }