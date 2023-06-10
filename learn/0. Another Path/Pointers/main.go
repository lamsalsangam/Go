package main

import "fmt"

// ################################################################

// INTRODUCTION TO POINTERS
// As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.

// x := 42
// // "x" is the name of a location in memory. That location is storing the integer value of 42
// A POINTER IS A VARIABLE
// A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.

// The * syntax defines a pointer:

// var p *int
// The & operator generates a pointer to its operand.

// myString := "hello"
// myStringPtr = &myString

// WHY ARE POINTERS USEFUL?
// Pointers allow us to manipulate data in memory directly, without making copies or duplicating data. This can make programs more efficient and allow us to do things that would be difficult or impossible without them.

// ===============================

// type Message struct {
// 	Recipient string
// 	Text      string
// }

// // Don't touch above this line

// func sendMessage(m Message) {
// 	// fmt.Printf("To: %v\n", &m.Recipient)
// 	// fmt.Printf("Message: %v\n", &m.Text)
// 	// Just dereference it
// 	fmt.Printf("To: %v\n", m.Recipient)
// 	fmt.Printf("Message: %v\n", m.Text)
// }

// // Don't touch below this line

// func test(recipient string, text string) {
// 	m := Message{Recipient: recipient, Text: text}
// 	sendMessage(m)
// 	fmt.Println("=====================================")
// }

// func main() {
// 	test("Lane", "Textio is getting better everyday!")
// 	test("Allan", "This pointer stuff is weird...")
// 	test("Tiffany", "What time will you be home for dinner?")
// }

// ################################################################
// POINTERS

// The * dereferences a pointer to gain access to the value

// fmt.Println(*myStringPtr) // read myString through the pointer
// *myStringPtr = "world"    // set myString through the pointer

// ===============================

// func removeProfanity(message *string) {
// 	// ?
// 	messageVal := *message
// 	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
// 	messageVal = strings.ReplaceAll(messageVal, "shoot", "****")
// 	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
// 	*message = messageVal
// }

// // don't touch below this line

// func test(messages []string) {
// 	for _, message := range messages {
// 		removeProfanity(&message)
// 		fmt.Println(message)
// 	}
// }

// func main() {
// 	messages1 := []string{
// 		"well shoot, this is awful",
// 		"dang robots",
// 		"dang them to heck",
// 	}

// 	messages2 := []string{
// 		"well shoot",
// 		"Allan is going straight to heck",
// 		"dang... that's a tough break",
// 	}

// 	test(messages1)
// 	test(messages2)
// }

// ################################################################

// NIL POINTERS
// Pointers can be very dangerous.

// If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

// ===============================

// func removeProfanity(message *string) {
// 	// ?
// 	if message == nil {
// 		return
// 	}
// 	messageVal := *message
// 	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
// 	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
// 	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
// 	*message = messageVal
// }

// // don't touch below this line

// func test(messages []string) {
// 	for _, message := range messages {
// 		if message == "" {
// 			removeProfanity(nil)
// 			fmt.Println("nil message detected")
// 		} else {
// 			removeProfanity(&message)
// 			fmt.Println(message)
// 		}
// 	}
// }

// func main() {
// 	messages := []string{
// 		"well shoot, this is awful",
// 		"",
// 		"dang robots",
// 		"dang them to heck",
// 		"",
// 	}

// 	messages2 := []string{
// 		"well shoot",
// 		"",
// 		"Allan is going straight to heck",
// 		"dang... that's a tough break",
// 		"",
// 	}

// 	test(messages)
// 	test(messages2)

// }

// ################################################################
// POINTER RECEIVERS
// A receiver type on a method can be a pointer.

// Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

// ===============================

// POINTER RECEIVER CODE
// Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

// type circle struct {
// 	x int
// 	y int
//     radius int
// }

// func (c *circle) grow(){
//     c.radius *= 2
// }

// func main(){
//     c := circle{
//         x: 1,
//         y: 2,
//         radius: 4,
//     }

//     // notice c is not a pointer in the calling function
//     // but the method still gains access to a pointer to c
//     c.grow()
//     fmt.Println(c.radius)
//     // prints 8
// }

// ===============================

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")

}
