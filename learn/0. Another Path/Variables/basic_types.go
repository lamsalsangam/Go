package main

import "fmt"

func main() {
	// Explicitly declared variables
	// var smsSendingLimit int
	// var costPerSMS float64
	// var hasPermission bool
	// var username string

	// Implicitly declared Variables
	smsSendingLimit := 0
	costPerSMS := 0.0
	var hasPermission = false
	username := ""

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
