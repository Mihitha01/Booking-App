package main

import (
	"fmt" // Handles formatted I/O, allowing us to print text to the console and format strings.
)

func main() {
	var conferenceName = "Go Conference" // var is used to declare changing varriables.
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking app!")
	fmt.Println("Remaining tickets:", remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. Remaining tickets: %v\n", userName, userTickets, remainingTickets)

}
