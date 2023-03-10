package main

import "fmt"

funct main(){
	msg := "Hello"
	writeMessage(&msg)
	fmt.Println(msg)
}

func writeMessage(msg *string) {
	*msg = "Hello Benita"
	fmt.Println(*msg)
}