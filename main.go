package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
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
