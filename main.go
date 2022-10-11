package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "MBS Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking platform \n", conferenceName)
	fmt.Printf("We have a total of %v and %v tickets are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var email string
		var ticketsOrdered uint
		fmt.Println("Enter your name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets to buy: ")
		fmt.Scan(&ticketsOrdered)

		isValidName := len(firstName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketRequest := ticketsOrdered <= remainingTickets
		isValidTicketInput := ticketsOrdered > 0

		if isValidName && isValidEmail && isValidTicketRequest && isValidTicketInput {
			remainingTickets = remainingTickets - ticketsOrdered
			bookings = append(bookings, firstName+",")

			fmt.Printf("First item in bookings is: %v. \n", bookings[0])
			fmt.Printf("Bookings is of type: %T. \n", bookings)
			fmt.Printf("We currently have %v booking(s). \n", len(bookings))

			fmt.Printf("Thank you %v for buying %v tickets \n", firstName, ticketsOrdered)
			fmt.Printf("We have %v tickets left \n", remainingTickets)
			fmt.Printf("The bookings are: %v. \n", bookings)
		} else {
			if !isValidName {
				fmt.Printf("%v is not a valid name\n", firstName)
			}
			if !isValidEmail {
				fmt.Printf("%v is not is not a valid email\n", email)
			}
			if !isValidTicketRequest {
				fmt.Printf("%v, We only have %v tickets left. So you can't book %v tickets\n",
					firstName, remainingTickets, ticketsOrdered)
			}
			if !isValidTicketInput {
				fmt.Printf("Tickets ordered should be greater than zero and not %v\n", ticketsOrdered)
			}
			continue
		}

		if remainingTickets == 0 {
			fmt.Printf("***** All tickets are booked, thank you! *****\n")
			break
		}
	}
}
