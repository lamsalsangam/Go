package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// ################################################################

// MUTEXES IN GO
// Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

// Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:

// .Lock()
// .Unlock()
// We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() method below.

// It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

// func protected(){
//     mux.Lock()
//     defer mux.Unlock()
//     // the rest of the function is protected
//     // any other calls to `mux.Lock()` will block
// }
// Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

// MAPS ARE NOT THREAD-SAFE
// Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.

// ===============================

// type safeCounter struct {
// 	counts map[string]int
// 	mux    *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounter) val(key string) int {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	return sc.counts[key]
// }

// // don't touch below this line

// func (sc safeCounter) slowIncrement(key string) {
// 	tempCounter := sc.counts[key]
// 	time.Sleep(time.Microsecond)
// 	tempCounter++
// 	sc.counts[key] = tempCounter
// }

// type emailTest struct {
// 	email string
// 	count int
// }

// func test(sc safeCounter, emailTests []emailTest) {
// 	emails := make(map[string]struct{})

// 	var wg sync.WaitGroup
// 	for _, emailT := range emailTests {
// 		emails[emailT.email] = struct{}{}
// 		for i := 0; i < emailT.count; i++ {
// 			wg.Add(1)
// 			go func(emailT emailTest) {
// 				sc.inc(emailT.email)
// 				wg.Done()
// 			}(emailT)
// 		}
// 	}
// 	wg.Wait()

// 	emailsSorted := make([]string, 0, len(emails))
// 	for email := range emails {
// 		emailsSorted = append(emailsSorted, email)
// 	}
// 	sort.Strings(emailsSorted)

// 	for _, email := range emailsSorted {
// 		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
// 	}
// 	fmt.Println("=====================================")
// }

// func main() {
// 	sc := safeCounter{
// 		counts: make(map[string]int),
// 		mux:    &sync.Mutex{},
// 	}
// 	test(sc, []emailTest{
// 		{
// 			email: "john@example.com",
// 			count: 23,
// 		},
// 		{
// 			email: "john@example.com",
// 			count: 29,
// 		},
// 		{
// 			email: "jill@example.com",
// 			count: 31,
// 		},
// 		{
// 			email: "jill@example.com",
// 			count: 67,
// 		},
// 	})
// 	test(sc, []emailTest{
// 		{
// 			email: "kaden@example.com",
// 			count: 23,
// 		},
// 		{
// 			email: "george@example.com",
// 			count: 126,
// 		},
// 		{
// 			email: "kaden@example.com",
// 			count: 31,
// 		},
// 		{
// 			email: "george@example.com",
// 			count: 453,
// 		},
// 	})
// }

// ################################################################

// RW MUTEX
// The standard library also exposes a sync.RWMutex

// In addition to these methods:

// Lock()
// Unlock()
// The sync.RWMutex also has these methods:

// RLock()
// RUnlock()
// The sync.RWMutex can help with performance if we have a read-intensive process. Many goroutines can safely read from the map at the same time (multiple Rlock() calls can happen simultaneously). However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.

// ===============================

type safeCounter struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

func test(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	sc.mux.RLock()
	defer sc.mux.RUnlock()
	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.RWMutex{},
	}
	test(sc, []emailTest{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	test(sc, []emailTest{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}
