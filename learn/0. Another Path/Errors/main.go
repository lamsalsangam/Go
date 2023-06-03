package main

import "fmt"

// ################################################################
// type error interface {
//     Error() string
// }

// // Atoi converts a stringified number to an interger
// i, err := strconv.Atoi("42b")
// if err != nil {
//     fmt.Println("couldn't convert:", err)
//     // because "42b" isn't a valid integer, we print:
//     // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
//     // Note:
//     // 'parsing "42b": invalid syntax' is returned by the .Error() method
//     return
// }
// // if we get here, then
// // i was converted successfully

// ===============================

// func main() {
// 	user, err := getUser()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	user, err := getUserProfile(user.ID)
// 	if (err) != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// func getUser()(User, error)  {

// }

// func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
// 	costForCustomer, err := sendSMS(msgToCustomer)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	costForSpouse, err := sendSMS(msgToSpouse)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	return costForCustomer + costForSpouse, nil
// }

// // don't edit below this line

// func sendSMS(message string) (float64, error) {
// 	const maxTextLen = 25
// 	const costPerChar = .0002
// 	if len(message) > maxTextLen {
// 		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
// 	}
// 	return costPerChar * float64(len(message)), nil
// }

// func test(msgToCustomer, msgToSpouse string) {
// 	defer fmt.Println("========")
// 	fmt.Println("Message for customer:", msgToCustomer)
// 	fmt.Println("Message for spouse:", msgToSpouse)
// 	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Total cost: $%.4f\n", totalCost)
// }

// func main() {
// 	test(
// 		"Thanks for coming in to our flower shop today!",
// 		"We hope you enjoyed your gift.",
// 	)
// 	test(
// 		"Thanks for joining us!",
// 		"Have a good day.",
// 	)
// 	test(
// 		"Thank you.",
// 		"Enjoy!",
// 	)
// 	test(
// 		"We loved having you in!",
// 		"We hope the rest of your evening is absolutely fantastic.",
// 	)
// }

// ################################################################

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent
	",cost,recipient)
}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}

// ################################################################