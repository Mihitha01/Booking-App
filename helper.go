package main

import "fmt"

// greetUsers encapsulates our introductory UI logic
func greetUsers(confName string, total int, remaining int) {
    fmt.Printf("Welcome to the %v booking application!\n", confName)
    fmt.Printf("We have a total of %v tickets, and %v are currently available.\n", total, remaining)
    fmt.Println("Follow the prompts to secure your spot.")
    fmt.Println("------------------------------------------------------")
}