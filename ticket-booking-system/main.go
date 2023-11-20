package main

import (
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	// var booking = [50]string{"Hello", "World", "!"}
	// var bookings [50]string
	// var bookings []string
	// bookings[0] = "Hello"
	// bookings[1] = "World!"
	// bookings[10] = "こんにちは"
	// bookings[11] = "世界!"

	for {

		firstName, lastName, email, userTickets := getuserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstnames()

			fmt.Printf("The first anmes of the bookings are: %v\n", firstNames)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			fmt.Println("#################################################")
			fmt.Println("#################################################")
			fmt.Println("#################################################")

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Hope to see you next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or the Last Name that you have entered is not true.")
			}
			if !isValidEmail {
				fmt.Println("Email doesn't contain the \"@\" sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Printf("Your input data is invalid. Try again")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v the conference booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		// firstNames = append(firstNames, names[0])
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

}

func getuserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email addtess")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for the user
	// var user = make(map[string]string)
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
