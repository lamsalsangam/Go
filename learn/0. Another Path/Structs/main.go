package main

import (
	"fmt"
)

// type messageToSend struct {
// 	phoneNumber int
// 	message     string
// }

// func test(m messageToSend) {
// 	fmt.Printf("Sendinf message: '%s' to: %v\n", m.message, m.phoneNumber)
// 	fmt.Println("=============================================")
// }

// func main() {
// 	test(messageToSend{
// 		phoneNumber: 1543135453543,
// 		message:     "Thanks for signing up",
// 	})
// 	test(messageToSend{
// 		phoneNumber: 4454654613463,
// 		message:     "Love to have you aboard",
// 	})
// }

// ################################################################
// Nested Structs

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.sender.name == "" || mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
// 		return false
// 	}
// 	return true
// }

// // don't touch below this line

// func test(mToSend messageToSend) {
// 	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
// 		mToSend.message,
// 		mToSend.sender.name,
// 		mToSend.sender.number,
// 		mToSend.recipient.name,
// 		mToSend.recipient.number,
// 	)
// 	fmt.Println()
// 	if canSendMessage(mToSend) {
// 		fmt.Println("...sent!")
// 	} else {
// 		fmt.Println("...can't send message")
// 	}
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(messageToSend{
// 		message: "you have an appointment tommorow",
// 		sender: user{
// 			name:   "Brenda Halafax",
// 			number: 16545550987,
// 		},
// 		recipient: user{
// 			name:   "Sally Sue",
// 			number: 19035558973,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have an event tommorow",
// 		sender: user{
// 			number: 16545550987,
// 		},
// 		recipient: user{
// 			name:   "Suzie Sall",
// 			number: 0,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have an party tommorow",
// 		sender: user{
// 			name:   "Njorn Halafax",
// 			number: 16545550987,
// 		},
// 		recipient: user{
// 			name:   "Sally Sue",
// 			number: 19035558973,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have a birthday tommorow",
// 		sender: user{
// 			name:   "Eli Halafax",
// 			number: 0,
// 		},
// 		recipient: user{
// 			name:   "Whitaker Sue",
// 			number: 19035558973,
// 		},
// 	})
// }

// ################################################################
// Anonymous structs
// It shall be inside the function to remove the error
// myCar := struct {
// 	Make string
// 	Model string
//   } {
// 	Make: "tesla",
// 	Model: "model 3"
//   }

// type car struct {
// 	Make   string
// 	Model  string
// 	Height int
// 	Width  int
// 	// Wheel is a field containing an anonymous struct
// 	Wheel struct {
// 		Radius   int
// 		Material string
// 	}
// }

// ################################################################
// Embedded Structs
// type car struct {
// 	make  string
// 	model string
// }

// type truck struct {
// 	car
// 	bedSize int
// }

// // This should be inside the function
// lanesTruck := truck{
// 	bedSize: 10,
// 	car: car{
// 		make: "Toyota",
// 		model: "camry",
// 	}
// }

// fmt.Println(lanesTruck.make)
// ===============================

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
