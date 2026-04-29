package main

import (
	"fmt" // Handles formatted I/O, allowing us to print text to the console and format strings.
)

func main() {
	var conferenceName = "Engineering Summit 2026" // var is used to declare changable variables.
	const totalTickets = 50
	var remainingTickets = 50
	var bookings []string

	greetUsers(conferenceName, totalTickets, remainingTickets)

	// 1. Defining variables with explicit types
	// When you declare a variable like this in Go,
	// the compiler automatically assigns it a "zero value."
	// For strings, this is an empty string (""), and for integers, it is 0.
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)

		fmt.Println("How many tickets would you like to book?")
		fmt.Scan(&userTickets)

		// 1. Data Validation Check
		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. You cannot book %v tickets.\n", remainingTickets, userTickets)
			fmt.Println("------------------------------------------------------")
			continue
		}

		// Update state and append to slice
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		// Output transaction summary
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
		fmt.Printf("Current list of bookings: %v\n", bookings)
		fmt.Println("------------------------------------------------------")

		// 2. Sold Out Check
		if remainingTickets == 0 {
			fmt.Println("Our conference is completely booked out. Thank you!")
			break
		}
	}

}
