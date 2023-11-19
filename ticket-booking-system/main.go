package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference Tickets is %T, remaining Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v the conference booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var booking = [50]string{"Hello", "World", "!"}
	var booking [50]string
	// booking[0] = "Hello"
	// booking[1] = "World!"
	// booking[10] = "こんにちは"
	// booking[11] = "世界!"

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

	remainingTickets = remainingTickets - userTickets
	booking[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", booking)
	fmt.Printf("The first value of the array: %v\n", booking[0])
	fmt.Printf("The whole array type: %T\n", booking)
	fmt.Printf("The whole array length: %v\n", len(booking))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
