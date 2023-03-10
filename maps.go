package main

import "fmt"

func main(){
//MAPS ARE A COLLECTION OF KEY VALUED PAIRS (they're a data structure in GO)
// However, when you create a map using a composite literal syntax, 
// the map is initialized with a default capacity of 0. 
// This means that the map can store some elements,
// but it might not be efficient for large numbers of elements,
// and may need to be resized dynamically during runtime, 
// which can cause performance issues.
// To avoid this problem, 
// you can use the make function to create a map with an initial capacity,
//  like this:

friends_colleagues := make(map[int][]string, 5 )
	friends_colleagues = map[int][]string{
		100 : {"Royce"},
		200 : {"Siri,", "Maro,", "Lucky"},
		300 : {"Daniella,", "Laju,", "David,", "Tegha,", "MJ"},
		400 : {"Ephraim,", "Elizabeth"},
		500 : {"Cynthia,","Yahaya"},
		}
		fmt.Println(friends_colleagues)
		fmt.Println("\n")
		fmt.Println(friends_colleagues[300])
		a := append(friends_colleagues[300], "Efe")
		friends_colleagues[600] = []string{"Okoh\n"}
		fmt.Println(friends_colleagues)
		fmt.Println("\n")
		fmt.Println(a)
//T check if key/map member is available, returns boolean value
		colleagues, ok := friends_colleagues[50]
		_ , ok_2 := friends_colleagues[700]//_ can also be used to tell the compiler that you're not using the variable
		fmt.Println(colleagues, ok)
		fmt.Println(colleagues, ok_2)

		fc := friends_colleagues
		fmt.Println(friends_colleagues)
		fmt.Println(fc)

		fc[100] = append(fc[100], "Roland\n")
// both maps fc[100] and friends_colleagues[100] are going to be affected, since it's always the memory location that is being passed (mem location of friends_colleagues not the elements)
		fmt.Println(friends_colleagues)
		fmt.Println(fc)

//To delete a key(member) of the map we use the keyword 'delete'
	fmt.Println("\n")
	friends_colleagues[700] =  []string{"Post-graduate students"} 
	fmt.Println(friends_colleagues)
	delete(friends_colleagues, 700 )
	fmt.Println(friends_colleagues)
	delete(friends_colleagues, 600)
	fmt.Println(friends_colleagues)


		
}