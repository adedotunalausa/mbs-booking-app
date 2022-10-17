package helper

import "strings"

func ValidateUserInput(firstName string, email string, ticketsOrdered uint, remainingTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketRequest := ticketsOrdered <= remainingTickets
	isValidTicketInput := ticketsOrdered > 0
	return isValidName, isValidEmail, isValidTicketRequest, isValidTicketInput
}
