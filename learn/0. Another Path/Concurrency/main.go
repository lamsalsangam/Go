package main

import (
	"fmt"
	"time"
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

// func addEmailsToQueue(emails []string) chan string {
// 	emailsToSend := make(chan string, len(emails))
// 	for _, email := range emails {
// 		emailsToSend <- email
// 	}
// 	return emailsToSend
// }

// // TEST SUITE - Don't Touch Below This Line

// func sendEmails(batchSize int, ch chan string) {
// 	for i := 0; i < batchSize; i++ {
// 		email := <-ch
// 		fmt.Println("Sending email:", email)
// 	}
// }

// func test(emails ...string) {
// 	fmt.Printf("Adding %v emails to queue...\n", len(emails))
// 	ch := addEmailsToQueue(emails)
// 	fmt.Println("Sending emails...")
// 	sendEmails(len(emails), ch)
// 	fmt.Println("==========================================")
// }

// func main() {
// 	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
// 	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
// 	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
// }

// ################################################################

// CLOSING CHANNELS IN GO
// Channels can be explicitly closed by a sender:

// ch := make(chan int)

// // do some stuff with the channel

// close(ch)
// CHECKING IF A CHANNEL IS CLOSED
// Similar to the ok value when accessing data in a map, receivers can check the ok value when receiving from a channel to test if a channel was closed.

// v, ok := <-ch
// ok is false if the channel is empty and closed.

// DON'T SEND ON A CLOSED CHANNEL
// Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

// Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

// ===============================

// func countReports(numSentCh chan int) int {
// 	// ?
// 	total := 0
// 	for {
// 		numSent, ok := <-numSentCh
// 		if !ok {
// 			break
// 		}
// 		total += numSent
// 	}
// 	return total
// }

// // don't touch below this line

// func test(numBatches int) {
// 	numSentCh := make(chan int)
// 	go sendReports(numBatches, numSentCh)

// 	fmt.Println("Start counting...")
// 	numReports := countReports(numSentCh)
// 	fmt.Printf("%v reports sent!\n", numReports)
// 	fmt.Println("========================")
// }

// func main() {
// 	test(3)
// 	test(4)
// 	test(5)
// 	test(6)
// }

// func sendReports(numBatches int, ch chan int) {
// 	for i := 0; i < numBatches; i++ {
// 		numReports := i*23 + 32%17
// 		ch <- numReports
// 		fmt.Printf("Sent batch of %v reports\n", numReports)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	close(ch)
// }

// ################################################################

// RANGE
// Similar to slices and maps, channels can be ranged over.

// for item := range ch {
//     // item is the next value received from the channel
// }
// This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

// ===============================

// func concurrrentFib(n int) {
// 	// ?
// 	chInts := make(chan int)
// 	go func() {
// 		fibonacci(n, chInts)
// 	}()
// 	for v := range chInts {
// 		fmt.Println(v)
// 	}
// }

// // don't touch below this line

// func test(n int) {
// 	fmt.Printf("Printing %v numbers...\n", n)
// 	concurrrentFib(n)
// 	fmt.Println("==============================")
// }

// func main() {
// 	test(10)
// 	test(5)
// 	test(20)
// 	test(13)
// }

// func fibonacci(n int, ch chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 		time.Sleep(time.Millisecond * 10)
// 	}
// 	close(ch)
// }

// ################################################################

// SELECT
// Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

// A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

// select {
//   case i, ok := <- chInts:
//     fmt.Println(i)
//   case s, ok := <- chStrings:
//     fmt.Println(s)
// }
// The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.

// ===============================

// func logMessages(chEmails, chSms chan string) {
// 	// ?
// 	for {
// 		select {
// 		case email, ok := <-chEmails:
// 			if !ok {
// 				return
// 			}
// 			logEmail(email)
// 		case sms, ok := <-chSms:
// 			if !ok {
// 				return
// 			}
// 			logSms(sms)
// 		}
// 	}
// }

// // don't touch below this line

// func logSms(sms string) {
// 	fmt.Println("SMS:", sms)
// }

// func logEmail(email string) {
// 	fmt.Println("Email:", email)
// }

// func test(sms []string, emails []string) {
// 	fmt.Println("Starting...")

// 	chSms, chEmails := sendToLogger(sms, emails)

// 	logMessages(chEmails, chSms)
// 	fmt.Println("===============================")
// }

// func main() {
// 	rand.Seed(0)
// 	test(
// 		[]string{
// 			"hi friend",
// 			"What's going on?",
// 			"Welcome to the business",
// 			"I'll pay you to be my friend",
// 		},
// 		[]string{
// 			"Will you make your appointment?",
// 			"Let's be friends",
// 			"What are you doing?",
// 			"I can't believe you've done this.",
// 		},
// 	)
// 	test(
// 		[]string{
// 			"this song slaps hard",
// 			"yooo hoooo",
// 			"i'm a big fan",
// 		},
// 		[]string{
// 			"What do you think of this song?",
// 			"I hate this band",
// 			"Can you believe this song?",
// 		},
// 	)
// }

// func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
// 	chSms = make(chan string)
// 	chEmails = make(chan string)
// 	go func() {
// 		for i := 0; i < len(sms) && i < len(emails); i++ {
// 			done := make(chan struct{})
// 			s := sms[i]
// 			e := emails[i]
// 			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			go func() {
// 				time.Sleep(t1)
// 				chSms <- s
// 				done <- struct{}{}
// 			}()
// 			go func() {
// 				time.Sleep(t2)
// 				chEmails <- e
// 				done <- struct{}{}
// 			}()
// 			<-done
// 			<-done
// 			time.Sleep(10 * time.Millisecond)
// 		}
// 		close(chSms)
// 		close(chEmails)
// 	}()
// 	return chSms, chEmails
// }

// ################################################################

// SELECT DEFAULT CASE
// The default case in a select statement executes immediately if no other channel has a value ready. A default case stops the select statement from blocking.

// select {
//   case v := <-ch:
//     // use v
//   default:
//     // receiving from ch would block
//     // so do something else
// }
// TICKERS
// time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
// time.After() sends a value once after the duration has passed.
// time.Sleep() blocks the current goroutine for the specified amount of time.
// READ-ONLY CHANNELS
// A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

// func main(){
//   ch := make(chan int)
//   readCh(ch)
// }

// func readCh(ch <-chan int) {
//   // ch can only be read from
//   // in this function
// }
// WRITE-ONLY CHANNELS
// The same goes for write-only channels, but the arrow's position moves.

// func writeCh(ch chan<- int) {
//   // ch can only be written to
//   // in this function
// }

// ===============================

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	// ?
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}

	}
}

// don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}
