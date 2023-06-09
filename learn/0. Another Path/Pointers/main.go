package main

import (
	"fmt"
	"strings"
)

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

func removeProfanity(message *string) {
	// ?
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
