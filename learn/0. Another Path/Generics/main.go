package main

import "fmt"

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

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}

// don't edit below this line

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount         int
	senderEmail    string
	recipientEmail string
}

func main() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")
}

func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}
