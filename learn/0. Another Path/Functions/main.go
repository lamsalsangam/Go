package main

import "fmt"

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

// func concat(s1, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	fmt.Println("XXXXXXX", " Happy Birthday")
// 	fmt.Println("Sangam", " Rich")
// 	fmt.Println("Go", " is Fantastic")
// }

// #################################################################

// func main() {
// 	sendsSoFar := 430
// 	const sendsToAdd = 25
// 	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
// 	fmt.Println("You've sent", sendsSoFar, "message")
// }

// func incrementSends(sendsSoFar, sendsToAdd int) int {
// 	sendsSoFar = sendsSoFar + sendsToAdd
// 	return sendsSoFar
// }

// #################################################################
// Named Return values
func yearsUntilEvents(age int) (
	yearsUntilAdult,
	yearsUntilDrinking,
	yearsUntilCarRental int) {

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	if yearsUntilDrinking = 21 - age; yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	if yearsUntilCarRental = 25 - age; yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func test(age int) {
	fmt.Println("Age", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("===================================================")
}

func main() {
	test(4)
	test(10)
	test(22)
}
