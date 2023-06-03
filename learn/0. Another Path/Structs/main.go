package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sendinf message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("=============================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 1543135453543,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 4454654613463,
		message:     "Love to have you aboard",
	})
}
