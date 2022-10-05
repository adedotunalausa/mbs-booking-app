package main

import "fmt"

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
		var ticketsOrdered uint
		fmt.Println("Enter your name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter the number of tickets to buy: ")
		fmt.Scan(&ticketsOrdered)

		if ticketsOrdered > remainingTickets {
			fmt.Printf("%v, We only have %v tickets left. So you can't book %v tickets\n",
				firstName, remainingTickets, ticketsOrdered)
			continue
		}

		remainingTickets = remainingTickets - ticketsOrdered
		bookings = append(bookings, firstName+",")

		fmt.Printf("First item in bookings is: %v. \n", bookings[0])
		fmt.Printf("Bookings is of type: %T. \n", bookings)
		fmt.Printf("We currently have %v booking(s). \n", len(bookings))

		fmt.Printf("Thank you %v for buying %v tickets \n", firstName, ticketsOrdered)
		fmt.Printf("We have %v tickets left \n", remainingTickets)
		fmt.Printf("The bookings are: %v. \n", bookings)

		if remainingTickets == 0 {
			fmt.Printf("")
		}
	}
}
