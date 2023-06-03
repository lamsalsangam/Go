package main

import (
	"fmt"
)

// ###In go the interface are implemented implicitly
// ################################################################
// Interface in Go
// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type rect struct {
// 	width, height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (r rect) perimeter() float64 {
// 	return 2*r.width + 2*r.height
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// ===============================

// func sendMessage(msg message) {
// 	fmt.Println(msg.getMessage())
// }

// type message interface {
// 	getMessage() string
// }

// // don't edit below this line

// type birthdayMessage struct {
// 	birthdayTime  time.Time
// 	recipientName string
// }

// func (bm birthdayMessage) getMessage() string {
// 	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
// }

// type sendingReport struct {
// 	reportName    string
// 	numberOfSends int
// }

// func (sr sendingReport) getMessage() string {
// 	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
// }

// func test(m message) {
// 	sendMessage(m)
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(sendingReport{
// 		reportName:    "First Report",
// 		numberOfSends: 10,
// 	})
// 	test(birthdayMessage{
// 		recipientName: "John Doe",
// 		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
// 	})
// 	test(sendingReport{
// 		reportName:    "First Report",
// 		numberOfSends: 10,
// 	})
// 	test(birthdayMessage{
// 		recipientName: "Bill Deer",
// 		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
// 	})
// }

// ################################################################
// Interface implemetation
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

// ?
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}
