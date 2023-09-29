package helper

import "strings"

func ValidateUser(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidNumber
}
