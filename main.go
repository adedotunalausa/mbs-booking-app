package main

import (
	"fmt"
	"strings"
)

var conferenceName = "MBS Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string

func main() {

	displayWelcomeMessages()

	for {
		firstName, lastName, email, ticketsOrdered := getUserInputs()

		isValidName, isValidEmail, isValidTicketRequest, isValidTicketInput :=
			validateUserInput(firstName, email, ticketsOrdered)

		if isValidName && isValidEmail && isValidTicketRequest && isValidTicketInput {
			bookTicket(firstName, lastName, ticketsOrdered)
		} else {
			handleErrors(isValidName, isValidEmail, isValidTicketRequest, isValidTicketInput,
				firstName, email, ticketsOrdered)
			continue
		}

		if remainingTickets == 0 {
			firstNames := getFirstNamesFromBookings()
			fmt.Printf("The following people have booked their tickets %v. \n", firstNames)
			fmt.Printf("***** All tickets are booked, thank you! *****\n")
			break
		}
	}
}

func displayWelcomeMessages() {
	fmt.Printf("Welcome to %v booking platform \n", conferenceName)
	fmt.Printf("We have a total of %v and %v tickets are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNamesFromBookings() []string {
	var firstnames []string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0]+",")
	}
	return firstnames
}

func validateUserInput(firstName string, email string, ticketsOrdered uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketRequest := ticketsOrdered <= remainingTickets
	isValidTicketInput := ticketsOrdered > 0
	return isValidName, isValidEmail, isValidTicketRequest, isValidTicketInput
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var ticketsOrdered uint

	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets to buy: ")
	fmt.Scan(&ticketsOrdered)
	return firstName, lastName, email, ticketsOrdered
}

func bookTicket(firstName string, lastName string, ticketsOrdered uint) {
	remainingTickets = remainingTickets - ticketsOrdered
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Bookings is of type: %T. \n", bookings)
	fmt.Printf("We currently have %v booking(s). \n", len(bookings))

	fmt.Printf("Thank you %v for buying %v tickets \n", firstName, ticketsOrdered)
	fmt.Printf("We have %v tickets left \n", remainingTickets)
}

func handleErrors(isValidName bool, isValidEmail bool, isValidTicketRequest bool, isValidTicketInput bool,
	firstName string, email string, ticketsOrdered uint) {
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
}
