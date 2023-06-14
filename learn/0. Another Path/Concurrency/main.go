package main

import (
	"fmt"
)

// ################################################################

// CONCURRENCY

// Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

// Concurrency is the ability to perform multiple tasks at the same time. Typically, our code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

// go doSomething()

// ===============================

// func sendEmail(message string) {
// 	// func() {
// 	go func() {
// 		time.Sleep(time.Millisecond * 250)
// 		fmt.Printf("Email received: '%s'\n", message)
// 	}()
// 	fmt.Printf("Email sent: '%s'\n", message)
// }

// // Don't touch below this line

// func test(message string) {
// 	sendEmail(message)
// 	time.Sleep(time.Millisecond * 500)
// 	fmt.Println("========================")
// }

// func main() {
// 	test("Hello there Stacy!")
// 	test("Hi there John!")
// 	test("Hey there Jane!")
// }

// ################################################################

// CHANNELS
// Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

// CREATE A CHANNEL
// Like maps and slices, channels must be created before use. They also use the same make keyword:

// ch := make(chan int)
// SEND DATA TO A CHANNEL
// ch <- 69
// The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.

// RECEIVE DATA FROM A CHANNEL
// v := <-ch
// This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.

// BLOCKING AND DEADLOCKS
// A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

// ===============================

// func filterOldEmails(emails []email) {
// 	isOldChan := make(chan bool)

// 	go sendIsOld(isOldChan, emails)

// 	isOld := <-isOldChan
// 	fmt.Println("email 1 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 2 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 3 is old:", isOld)
// }

// // TEST SUITE -- Don't touch below this line

// func sendIsOld(isOldChan chan<- bool, emails []email) {
// 	for _, e := range emails {
// 		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
// 			isOldChan <- true
// 			continue
// 		}
// 		isOldChan <- false
// 	}
// }

// type email struct {
// 	body string
// 	date time.Time
// }

// func test(emails []email) {
// 	filterOldEmails(emails)
// 	fmt.Println("==========================================")
// }

// func main() {
// 	test([]email{
// 		{
// 			body: "Are you going to make it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "I need a break",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What were you thinking?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Yo are you okay?",
// 			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Have you heard of that website Boot.dev?",
// 			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "It's awesome honestly.",
// 			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Today is the day!",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What do you want for lunch?",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Why are you the way that you are?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Did we do it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Letsa Go!",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Okay...?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// }

// ################################################################
// More on CHANNELS
// Empty structs are often used as tokens in Go programs. In this context, a token is a unary value. In other words, we don't care what is passed through the channel. We care when and if it is passed.

// We can block and wait until something is sent on a channel using the following syntax

// <-ch
// This will block until it pops a single item off the channel, then continue, discarding the item.

// ===============================

// func waitForDbs(numDBs int, dbChan chan struct{}) {
// 	// ?
// 	for i := 0; i < numDBs; i++ {
// 		<-dbChan
// 	}
// }

// // don't touch below this line

// func test(numDBs int) {
// 	dbChan := getDatabasesChannel(numDBs)
// 	fmt.Printf("Waiting for %v databases...\n", numDBs)
// 	waitForDbs(numDBs, dbChan)
// 	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
// 	fmt.Println("All databases are online!")
// 	fmt.Println("=====================================")
// }

// func main() {
// 	test(3)
// 	test(4)
// 	test(5)
// }

// func getDatabasesChannel(numDBs int) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		for i := 0; i < numDBs; i++ {
// 			ch <- struct{}{}
// 			fmt.Printf("Database %v is online\n", i+1)
// 		}
// 	}()
// 	return ch
// }

// ################################################################
// BUFFERED CHANNELS
// Channels can optionally be buffered.

// CREATING A CHANNEL WITH A BUFFER
// You can provide a buffer length as the second argument to make() to create a buffered channel:

// ch := make(chan int, 100)
// Sending on a buffered channel only blocks when the buffer is full.

// Receiving blocks only when the buffer is empty.

// ===============================

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}
