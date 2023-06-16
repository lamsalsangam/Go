package main

import (
	"errors"
	"fmt"
	"time"
)

// ################################################################

// GENERICS IN GO
// As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many circumstances. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't really care about the values stored in the slice. Unfortunately in Go we would need to write it multiple times for each type, which is a very un-DRY thing to do.

// func splitIntSlice(s []int) ([]int, []int) {
//     mid := len(s)/2
//     return s[:mid], s[mid:]
// }
// func splitStringSlice(s []string) ([]string, []string) {
//     mid := len(s)/2
//     return s[:mid], s[mid:]
// }
// In Go 1.20 however, support for generics was released, effectively solving this problem!

// TYPE PARAMETERS
// Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

// func splitAnySlice[T any](s []T) ([]T, []T) {
//     mid := len(s)/2
//     return s[:mid], s[mid:]
// }
// In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.

// firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
// fmt.Println(firstInts, secondInts)

// ===============================

// func getLast[T any](s []T) T {
// 	if len(s) == 0 {
// 		var zeroVal T
// 		return zeroVal
// 	}
// 	return s[len(s)-1]
// }

// // don't edit below this line

// type email struct {
// 	message        string
// 	senderEmail    string
// 	recipientEmail string
// }

// type payment struct {
// 	amount         int
// 	senderEmail    string
// 	recipientEmail string
// }

// func main() {
// 	test([]email{}, "email")
// 	test([]email{
// 		{
// 			"Hi Margo",
// 			"janet@example.com",
// 			"margo@example.com",
// 		},
// 		{
// 			"Hey Margo I really wanna chat",
// 			"janet@example.com",
// 			"margo@example.com",
// 		},
// 		{
// 			"ANSWER ME",
// 			"janet@example.com",
// 			"margo@example.com",
// 		},
// 	}, "email")
// 	test([]payment{
// 		{
// 			5,
// 			"jane@example.com",
// 			"sally@example.com",
// 		},
// 		{
// 			25,
// 			"jane@example.com",
// 			"mark@example.com",
// 		},
// 		{
// 			1,
// 			"jane@example.com",
// 			"sally@example.com",
// 		},
// 		{
// 			16,
// 			"jane@example.com",
// 			"margo@example.com",
// 		},
// 	}, "payment")
// }

// func test[T any](s []T, desc string) {
// 	last := getLast(s)
// 	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
// 	for i, v := range s {
// 		fmt.Printf("Item #%v: %v\n", i+1, v)
// 	}
// 	fmt.Printf("Last item in list: %v\n", last)
// 	fmt.Println(" --- ")
// }

// ################################################################

// CONSTRAINTS
// Sometimes you need the logic in your generic function to know something about the types it operates on. The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

// func splitAnySlice[T any](s []T) ([]T, []T) {
//     mid := len(s)/2
//     return s[:mid], s[mid:]
// }
// Constraints are just interfaces that allow us to write generics that only operate within the constraint of a given interface type. In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.

// CREATING A CUSTOM CONSTRAINT
// Let's take a look at the example of a concat function. It takes a slice of values and concatenates the values into a string. This should work with any type that can represent itself as a string, even if it's not a string under the hood. For example, a user struct can have a .String() that returns a string with the user's name and age.

// type stringer interface {
//     String() string
// }

// func concat[T stringer](vals []T) string {
//     result := ""
//     for _, val := range vals {
//         // this is where the .String() method
//         // is used. That's why we need a more specific
//         // constraint instead of the any constraint
//         result += val.String()
//     }
//     return result
// }

// ===============================

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	// ?
	newBalance := balance - newItem.GetCost()
	if newBalance < 0.0 {
		return nil, 0.0, errors.New("Insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	test(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	test(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)
}

func test[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
}
