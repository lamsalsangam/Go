package main

import (
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
// func getLogger(formatter func(string, string) string) func(string, string) {
// 	// ?
// 	return func(s1, s2 string) {
// 		fmt.Println(formatter(s1, s2))
// 	}
// }

// // don't touch below this line

// func test(first string, errors []error, formatter func(string, string) string) {
// 	defer fmt.Println("====================================")
// 	logger := getLogger(formatter)
// 	fmt.Println("Logs:")
// 	for _, err := range errors {
// 		logger(first, err.Error())
// 	}
// }

// func colonDelimit(first, second string) string {
// 	return first + ": " + second
// }
// func commaDelimit(first, second string) string {
// 	return first + ", " + second
// }

// func main() {
// 	dbErrors := []error{
// 		errors.New("out of memory"),
// 		errors.New("cpu is pegged"),
// 		errors.New("networking issue"),
// 		errors.New("invalid syntax"),
// 	}
// 	test("Error on database server", dbErrors, colonDelimit)

// 	mailErrors := []error{
// 		errors.New("email too large"),
// 		errors.New("non alphanumeric symbols found"),
// 	}
// 	test("Error on mail server", mailErrors, commaDelimit)
// }

// ################################################################
// DEFER
// The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

// ===============================

// const (
// 	logDeleted  = "user deleted"
// 	logNotFound = "user not found"
// 	logAdmin    = "admin deleted"
// )

// func logAndDelete(users map[string]user, name string) (log string) {
// 	defer delete(users, name)
// 	user, ok := users[name]
// 	if !ok {
// 		delete(users, name)
// 		return logNotFound
// 	}
// 	if user.admin {
// 		return logAdmin
// 	}
// 	delete(users, name)
// 	return logDeleted
// }

// // don't touch below this line

// type user struct {
// 	name   string
// 	number int
// 	admin  bool
// }

// func test(users map[string]user, name string) {
// 	fmt.Printf("Attempting to delete %s...\n", name)
// 	defer fmt.Println("====================================")
// 	log := logAndDelete(users, name)
// 	fmt.Println("Log:", log)
// }

// func main() {
// 	users := map[string]user{
// 		"john": {
// 			name:   "john",
// 			number: 18965554631,
// 			admin:  true,
// 		},
// 		"elon": {
// 			name:   "elon",
// 			number: 19875556452,
// 			admin:  true,
// 		},
// 		"breanna": {
// 			name:   "breanna",
// 			number: 98575554231,
// 			admin:  false,
// 		},
// 		"kade": {
// 			name:   "kade",
// 			number: 10765557221,
// 			admin:  false,
// 		},
// 	}

// 	fmt.Println("Initial users:")
// 	usersSorted := []string{}
// 	for name := range users {
// 		usersSorted = append(usersSorted, name)
// 	}
// 	sort.Strings(usersSorted)
// 	for _, name := range usersSorted {
// 		fmt.Println(" -", name)
// 	}
// 	fmt.Println("====================================")

// 	test(users, "john")
// 	test(users, "santa")
// 	test(users, "kade")

// 	fmt.Println("Final users:")
// 	usersSorted = []string{}
// 	for name := range users {
// 		usersSorted = append(usersSorted, name)
// 	}
// 	sort.Strings(usersSorted)
// 	for _, name := range usersSorted {
// 		fmt.Println(" -", name)
// 	}
// 	fmt.Println("====================================")
// }

// ################################################################
// CLOSURES
// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

// In this example, the concatter() function returns a function that has reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates that same doc variable.

// func concatter() func(string) string {
// 	doc := ""
// 	return func(word string) string {
// 		doc += word + " "
// 		return doc
// 	}
// }

// func main() {
// 	harryPotterAggregator := concatter()
// 	harryPotterAggregator("Mr.")
// 	harryPotterAggregator("and")
// 	harryPotterAggregator("Mrs.")
// 	harryPotterAggregator("Dursley")
// 	harryPotterAggregator("of")
// 	harryPotterAggregator("number")
// 	harryPotterAggregator("four,")
// 	harryPotterAggregator("Privet")

// 	fmt.Println(harryPotterAggregator("Drive"))
// 	// Mr. and Mrs. Dursley of number four, Privet Drive
// }

// ===============================

func adder() func(int) int {
	// ?
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})
}
