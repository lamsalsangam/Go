package main

import (
	"errors"
	"fmt"
)

// ################################################################

// FIRST CLASS AND HIGHER ORDER FUNCTIONS
// A programming language is said to have "first-class functions" when functions in that language are treated like any other variable. For example, in such a language, a function can be passed as an argument to other functions, can be returned by another function and can be assigned as a value to a variable.

// A function that returns a function or accepts a function as input is called a Higher-Order Function.

// Go supports first-class and higher-order functions. Another way to think of this is that a function is just another type -- just like ints and strings and bools.

// For example, to accept a function as a parameter:

// func add(x, y int) int {
//   return x + y
// }

// func mul(x, y int) int {
//   return x * y
// }

// // aggregate applies the given math function to the first 3 inputs
// func aggregate(a, b, c int, arithmetic func(int, int) int) int {
//   return arithmetic(arithmetic(a, b), c)
// }

// func main(){
//   fmt.Println(aggregate(2,3,4, add))
//   // prints 9
//   fmt.Println(aggregate(2,3,4, mul))
//   // prints 24
// }

// ===============================

// func getFormattedMessages(messages []string, formatter func(string) string) []string {
// 	formattedMessages := []string{}
// 	for _, message := range messages {
// 		formattedMessages = append(formattedMessages, formatter(message))
// 	}
// 	return formattedMessages
// }

// // don't touch below this line

// func addSignature(message string) string {
// 	return message + " Kind regards."
// }

// func addGreeting(message string) string {
// 	return "Hello! " + message
// }

// func test(messages []string, formatter func(string) string) {
// 	defer fmt.Println("====================================")
// 	formattedMessages := getFormattedMessages(messages, formatter)
// 	if len(formattedMessages) != len(messages) {
// 		fmt.Println("The number of messages returned is incorrect.")
// 		return
// 	}
// 	for i, message := range messages {
// 		formatted := formattedMessages[i]
// 		fmt.Printf(" * %s -> %s\n", message, formatted)
// 	}
// }

// func main() {
// 	test([]string{
// 		"Thanks for getting back to me.",
// 		"Great to see you again.",
// 		"I would love to hang out this weekend.",
// 		"Got any hot stock tips?",
// 	}, addSignature)
// 	test([]string{
// 		"Thanks for getting back to me.",
// 		"Great to see you again.",
// 		"I would love to hang out this weekend.",
// 		"Got any hot stock tips?",
// 	}, addGreeting)
// }

// ################################################################
// CURRYING
// Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.

// ===============================

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	// ?
	return func(s1, s2 string) {
		fmt.Println(formatter(s1, s2))
	}
}

// don't touch below this line

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}
