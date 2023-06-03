package main

import "fmt"

func main() {
	// Explicitly declared variables
	// var smsSendingLimit int
	// var costPerSMS float64
	// var hasPermission bool
	// var username string

	// Implicitly declared Variables
	// smsSendingLimit := 0
	// costPerSMS := 0.0
	// var hasPermission = false
	// username := ""

	// fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// #################################################################
	// Same Line Declaration
	// averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	// fmt.Println(averageOpenRate, displayMessage)

	// #################################################################
	// Type Sizes
	// accountAge := 2.6
	// accountAgeInt := int(accountAge)
	// fmt.Println("Your account has existed for", accountAgeInt, "years")

	// #################################################################
	// Constants
	// const premiumPlanName = "Premium Plan"
	// const basicPlanName = "Basic Plan"

	// fmt.Println("plan:", premiumPlanName)
	// fmt.Println("plan:", basicPlanName)

	// #################################################################
	// Conditionals
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not sent")
	}
}
