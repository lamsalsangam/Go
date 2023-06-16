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

// func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
// 	// ?
// 	newBalance := balance - newItem.GetCost()
// 	if newBalance < 0.0 {
// 		return nil, 0.0, errors.New("Insufficient funds")
// 	}
// 	oldItems = append(oldItems, newItem)
// 	return oldItems, newBalance, nil
// }

// // don't edit below this line

// type lineItem interface {
// 	GetCost() float64
// 	GetName() string
// }

// type subscription struct {
// 	userEmail string
// 	startDate time.Time
// 	interval  string
// }

// func (s subscription) GetName() string {
// 	return fmt.Sprintf("%s subscription", s.interval)
// }

// func (s subscription) GetCost() float64 {
// 	if s.interval == "monthly" {
// 		return 25.00
// 	}
// 	if s.interval == "yearly" {
// 		return 250.00
// 	}
// 	return 0.0
// }

// type oneTimeUsagePlan struct {
// 	userEmail        string
// 	numEmailsAllowed int
// }

// func (otup oneTimeUsagePlan) GetName() string {
// 	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
// }

// func (otup oneTimeUsagePlan) GetCost() float64 {
// 	const costPerEmail = 0.03
// 	return float64(otup.numEmailsAllowed) * costPerEmail
// }

// func main() {
// 	test(subscription{
// 		userEmail: "john@example.com",
// 		startDate: time.Now().UTC(),
// 		interval:  "yearly",
// 	},
// 		[]subscription{},
// 		1000.00,
// 	)
// 	test(subscription{
// 		userEmail: "jane@example.com",
// 		startDate: time.Now().UTC(),
// 		interval:  "monthly",
// 	},
// 		[]subscription{
// 			{
// 				userEmail: "jane@example.com",
// 				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
// 				interval:  "monthly",
// 			},
// 			{
// 				userEmail: "jane@example.com",
// 				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
// 				interval:  "yearly",
// 			},
// 		},
// 		686.20,
// 	)
// 	test(oneTimeUsagePlan{
// 		userEmail:        "dillon@example.com",
// 		numEmailsAllowed: 5000,
// 	},
// 		[]oneTimeUsagePlan{},
// 		756.20,
// 	)
// 	test(oneTimeUsagePlan{
// 		userEmail:        "dalton@example.com",
// 		numEmailsAllowed: 100000,
// 	},
// 		[]oneTimeUsagePlan{
// 			{
// 				userEmail:        "dalton@example.com",
// 				numEmailsAllowed: 34200,
// 			},
// 		},
// 		32.20,
// 	)
// }

// func test[T lineItem](newItem T, oldItems []T, balance float64) {
// 	fmt.Println(" --- ")
// 	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
// 	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
// 	if err != nil {
// 		fmt.Printf("Got error: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
// }

// ################################################################

// PARAMETRIC CONSTRAINTS
// Your interface definitions, which can later be used as constraints, can accept type parameters as well.

// // The store interface represents a store that sells products.
// // It takes a type parameter P that represents the type of products the store sells.
// type store[P product] interface {
// 	Sell(P)
// }

// type product interface {
// 	Price() float64
// 	Name() string
// }

// type book struct {
// 	title  string
// 	author string
// 	price  float64
// }

// func (b book) Price() float64 {
// 	return b.price
// }

// func (b book) Name() string {
// 	return fmt.Sprintf("%s by %s", b.title, b.author)
// }

// type toy struct {
// 	name  string
// 	price float64
// }

// func (t toy) Price() float64 {
// 	return t.price
// }

// func (t toy) Name() string {
// 	return t.name
// }

// // The bookStore struct represents a store that sells books.
// type bookStore struct {
// 	booksSold []book
// }

// // Sell adds a book to the bookStore's inventory.
// func (bs *bookStore) Sell(b book) {
// 	bs.booksSold = append(bs.booksSold, b)
// }

// // The toyStore struct represents a store that sells toys.
// type toyStore struct {
// 	toysSold []toy
// }

// // Sell adds a toy to the toyStore's inventory.
// func (ts *toyStore) Sell(t toy) {
// 	ts.toysSold = append(ts.toysSold, t)
// }

// // sellProducts takes a store and a slice of products and sells
// // each product one by one.
// func sellProducts[P product](s store[P], products []P) {
// 	for _, p := range products {
// 		s.Sell(p)
// 	}
// }

// func main() {
// 	bs := bookStore{
// 		booksSold: []book{},
// 	}

//     // By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
// 	sellProducts[book](&bs, []book{
// 		{
// 			title:  "The Hobbit",
// 			author: "J.R.R. Tolkien",
// 			price:  10.0,
// 		},
// 		{
// 			title:  "The Lord of the Rings",
// 			author: "J.R.R. Tolkien",
// 			price:  20.0,
// 		},
// 	})
// 	fmt.Println(bs.booksSold)

//     // We can then do the same for toys
// 	ts := toyStore{
// 		toysSold: []toy{},
// 	}
// 	sellProducts[toy](&ts, []toy{
// 		{
// 			name:  "Lego",
// 			price: 10.0,
// 		},
// 		{
// 			name:  "Barbie",
// 			price: 20.0,
// 		},
// 	})
// 	fmt.Println(ts.toysSold)
// }

// ===============================

// ?
type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "joe@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "samuel.boggs@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "pro"},
		user{UserEmail: "jade.row@example.com"},
	)
	testBiller[org](
		orgBiller{Plan: "basic"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
	testBiller[org](
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
}

func testBiller[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created for %v dollars\n", bill.Amount)
	fmt.Println(" --- ")
}
