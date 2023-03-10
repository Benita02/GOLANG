package main
//Code to input values of 1 - 100 into an array x
import "fmt"

func main(){
	var x [100]int
	count := 0
	x[0] = 1
	i := 0
	for i < 100 {
		
		x[i] = count
		//fmt.Println(x[i])
		//fmt.Println("\n")
		i = i + 1
		count = count + 1
	}
	fmt.Println(x[63])

}