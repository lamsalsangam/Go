package helper

import "strings"

// To export the function to the other packages we need to capitalize the name of the function.
func ValidateUserInput(firstName string, lastName string, remainingTickets uint, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
