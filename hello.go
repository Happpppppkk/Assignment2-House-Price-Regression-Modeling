package main

import (
	"fmt"
	"strconv"
	"week1/helper"
)

// package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidNumber := helper.ValidateUser(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidNumber {
			bookTicket(userTickets, firstName, lastName, email)
			//function print user firstname
			firstNames := getsFirstName()
			fmt.Printf("The first name of bookings are %v", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end the program
				fmt.Println("our conference is Booked out, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("name is too short")
			}
			if !isValidEmail {
				fmt.Println("email is wrong")
			}
			if !isValidNumber {
				fmt.Println("bad data entry, try again")
			}

		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v seats and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket and attend!")
}

func getsFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings { //iterate list of string or map
		//var names = strings.Fields(booking)

		firstNames = append(firstNames, booking["firstName"]) //"firstName" is key name

	}
	return firstNames

}

// get user input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter your email")
	fmt.Scan(&email)

	fmt.Println("enter the number of tickets you want")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// booking application function
func bookTicket(userTickets uint, firstName string,
	lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	//var userData = make(map[string]string)
	var userData = make(map[string]string) //define type of the map w/[string]string
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	//converse userTickets from uint to string
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10) //10 stand for decimal number

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("user %v %v book %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remain for this conference\n", remainingTickets)

	//return bookings

}
