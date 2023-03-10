package main

import "fmt"

func main() {
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
		fmt.Println(identityMatrix)
		fmt.Println("\n")
		identityMatrix[0][1] = 2
		identityMatrix[0][2] = 3
		identityMatrix[1][2] = 2
		identityMatrix[2][2] = 0
		fmt.Println(identityMatrix) 
		 
		arr2 := identityMatrix[2:5]
		fmt.Printf{"%v\n", arr2}
}