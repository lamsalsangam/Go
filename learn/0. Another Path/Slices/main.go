package main

import (
	"fmt"
)

// To declare an array of 10 integers:

// var myInts [10]int

// or to declare an initialized literal:

// primes := [6]int{2, 3, 5, 7, 11, 13}

// to declare the slices

// mySlice := primes[1:4]

// mySlice := []string{"I", "love", "go"}

// ===============================

// const (
// 	retry1 = "click here to sign up"
// 	retry2 = "pretty please click here"
// 	retry3 = "we beg you to sign up"
// )

// func getMessageWithRetries() [3]string {
// 	// ?
// 	return [3]string{
// 		retry1,
// 		retry2,
// 		retry3,
// 	}
// }

// // don't touch below this line

// func testSend(name string, doneAt int) {
// 	fmt.Printf("sending to %v...", name)
// 	fmt.Println()

// 	messages := getMessageWithRetries()
// 	for i := 0; i < len(messages); i++ {
// 		msg := messages[i]
// 		fmt.Printf(`sending: "%v"`, msg)
// 		fmt.Println()
// 		if i == doneAt {
// 			fmt.Println("they responded!")
// 			break
// 		}
// 		if i == len(messages)-1 {
// 			fmt.Println("complete failure")
// 		}
// 	}
// }

// func main() {
// 	testSend("Bob", 0)
// 	testSend("Alice", 1)
// 	testSend("Mangalam", 2)
// 	testSend("Ozgur", 3)
// }

// ################################################################
// Slices Review

// const (
// 	planFree = "free"
// 	planPro  = "pro"
// )

// func getMessageWithRetriesForPlan(plan string) ([]string, error) {
// 	allMessages := getMessageWithRetries()
// 	// ?
// 	if plan == planPro {
// 		return allMessages[:], nil
// 	}
// 	if plan == planFree {
// 		return allMessages[:2], nil
// 	}
// 	return nil, errors.New("unsupported plan")
// }

// // don't touch below this line

// func getMessageWithRetries() [3]string {
// 	return [3]string{
// 		"click here to sign up",
// 		"pretty please click here",
// 		"we beg you to sign up",
// 	}
// }

// func test(name string, doneAt int, plan string) {
// 	defer fmt.Println("=====================================")
// 	fmt.Printf("sending to %v...", name)
// 	fmt.Println()

// 	messages, err := getMessageWithRetriesForPlan(plan)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	for i := 0; i < len(messages); i++ {
// 		msg := messages[i]
// 		fmt.Printf(`sending: "%v"`, msg)
// 		fmt.Println()
// 		if i == doneAt {
// 			fmt.Println("they responded!")
// 			break
// 		}
// 		if i == len(messages)-1 {
// 			fmt.Println("no response")
// 		}
// 	}
// }

// func main() {
// 	test("Ozgur", 3, planFree)
// 	test("Jeff", 3, planPro)
// 	test("Sally", 2, planPro)
// 	test("Sally", 3, "no plan")
// }

// ################################################################
// Make the slice

// func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
// mySlice := make([]int, 5)

// ===============================

// func getMessageCosts(messages []string) []float64 {
// 	// ?
// 	costs := make([]float64, len(messages))
// 	for i := 0; i < len(messages); i++ {
// 		message := messages[i]
// 		cost := float64(len(message)) * 0.01
// 		costs[i] = cost
// 	}
// 	return costs
// }

// // don't edit below this line

// func test(messages []string) {
// 	costs := getMessageCosts(messages)
// 	fmt.Println("Messages:")
// 	for i := 0; i < len(messages); i++ {
// 		fmt.Printf(" - %v\n", messages[i])
// 	}
// 	fmt.Println("Costs:")
// 	for i := 0; i < len(costs); i++ {
// 		fmt.Printf(" - %.2f\n", costs[i])
// 	}
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test([]string{
// 		"Welcome to the movies!",
// 		"Enjoy your popcorn!",
// 		"Please don't talk during the movie!",
// 	})
// 	test([]string{
// 		"I don't want to be here anymore",
// 		"Can we go home?",
// 		"I'm hungry",
// 		"I'm bored",
// 	})
// 	test([]string{
// 		"Hello",
// 		"Hi",
// 		"Hey",
// 		"Hi there",
// 		"Hey there",
// 		"Hi there",
// 		"Hello there",
// 		"Hey there",
// 		"Hello there",
// 		"General Kenobi",
// 	})
// }

// ################################################################
// Variadic

// Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

// A variadic function receives the variadic arguments as a slice.

// func concat(strs ...string) string {
//     final := ""
//     // strs is just a slice of strings
//     for _, str := range strs {
//         final += str
//     }
//     return final
// }

// func main() {
//     final := concat("Hello ", "there ", "friend!")
//     fmt.Println(total)
//     // Output: Hello there friend!
// }

// ===============================

// func sum(nums ...float64) float64 {
// 	// ?
// 	total := 0.0
// 	for i := 0; i < len(nums); i++ {
// 		total += nums[i]
// 	}
// 	return total
// }

// // don't edit below this line

// func test(nums ...float64) {
// 	total := sum(nums...)
// 	fmt.Printf("Summing %v costs...\n", len(nums))
// 	fmt.Printf("Bill for the month: %.2f\n", total)
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test(1.0, 2.0, 3.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
// }

// ################################################################
// Append

// The built-in append function is used to dynamically add elements to a slice:

// func append(slice []Type, elems ...Type) []Type
// If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.

// Notice that append() is variadic, the following are all valid:

// slice = append(slice, oneThing)
// slice = append(slice, firstThing, secondThing)
// slice = append(slice, anotherSlice...)

// ===============================

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// ?
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}

// ################################################################
