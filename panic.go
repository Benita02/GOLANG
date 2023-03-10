package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Starting program...")

    panic("Something went wrong!")

    fmt.Println("Program completed.")
}