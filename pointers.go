package main

import "fmt"

type Company_you_keep struct {
	characters []string
	locations []string
	rating float32
}
func main(){
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(*b)



/// Declare a pointer variable company_you_keep of type *Company_you_keep
var company_you_keep *Company_you_keep

// Use new() to allocate memory for a new Company_you_keep struct
company_you_keep = new(Company_you_keep)

// Print the pointer variable company_you_keep, which should now hold the memory address of the new struct
fmt.Println(company_you_keep)

// Access the rating field of the struct using a dereference operator
fmt.Println(company_you_keep.rating)

// Set the value of the rating field
(*company_you_keep).rating = 8.0

// Access and modify the characters field of the struct
(*company_you_keep).characters = []string{"Birdie", "David", "Daphne", "Ollie"}

//Go compiler automatically dereferences for us incase we forget
(company_you_keep).locations = []string{"America", "Ireland"}
// Print the rating and characters fields
fmt.Println((*company_you_keep).rating)
fmt.Println((company_you_keep).locations)//code still runs without `*` dereference operator
fmt.Println((*company_you_keep).characters)
}