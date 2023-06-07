package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

// ################################################################

// MAPS
// Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

// The zero value of a map is nil.

// We can create a map by using a literal or by using the make() function:

// ages := make(map[string]int)
// ages["John"] = 37
// ages["Mary"] = 24
// ages["Mary"] = 21 // overwrites 24
// ages = map[string]int{
//   "John": 37,
//   "Mary": 21,
// }
// The len() function works on a map, it returns the total number of key/value pairs.

// ages = map[string]int{
//   "John": 37,
//   "Mary": 21,
// }
// fmt.Println(len(ages)) // 2

// ===============================

// func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
// 	// ?
// 	userMap := make(map[string]user)
// 	if len(names) != len(phoneNumbers) {
// 		return nil, errors.New("invalid sizes")
// 	}
// 	for i := 0; i < len(names); i++ {
// 		name := names[i]
// 		phoneNumber := phoneNumbers[i]

// 		userMap[name] = user{name: name, phoneNumber: phoneNumber}
// 	}
// 	return userMap, nil
// }

// // don't touch below this line

// type user struct {
// 	name        string
// 	phoneNumber int
// }

// func test(names []string, phoneNumbers []int) {
// 	fmt.Println("Creating map...")
// 	defer fmt.Println("====================================")
// 	users, err := getUserMap(names, phoneNumbers)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, name := range names {
// 		fmt.Printf("key: %v, value:\n", name)
// 		fmt.Println(" - name:", users[name].name)
// 		fmt.Println(" - number:", users[name].phoneNumber)
// 	}
// }

// func main() {
// 	test(
// 		[]string{"John", "Bob", "Jill"},
// 		[]int{14355550987, 98765550987, 18265554567},
// 	)
// 	test(
// 		[]string{"John", "Bob"},
// 		[]int{14355550987, 98765550987, 18265554567},
// 	)
// 	test(
// 		[]string{"George", "Sally", "Rich", "Sue"},
// 		[]int{20955559812, 38385550982, 48265554567, 16045559873},
// 	)
// }

// ################################################################

// // MUTATIONS
// // INSERT AN ELEMENT
// m[key] = elem
// // GET AN ELEMENT
// elem = m[key]
// // DELETE AN ELEMENT
// delete(m, key)
// // CHECK IF A KEY EXISTS
// elem, ok := m[key]
// // If key is in m, then ok is true. If not, ok is false.

// // If key is not in the map, then elem is the zero value for the map's element type.

// ===============================

// func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
// 	// ?
// 	existingUser, ok := users[name]
// 	if !ok {
// 		return false, errors.New("not found")
// 	}
// 	if existingUser.scheduledForDeletion {
// 		delete(users, name)
// 		return true, nil
// 	}
// 	return false, nil
// }

// // don't touch below this line

// type user struct {
// 	name                 string
// 	number               int
// 	scheduledForDeletion bool
// }

// func test(users map[string]user, name string) {
// 	fmt.Printf("Attempting to delete %s...\n", name)
// 	defer fmt.Println("====================================")
// 	deleted, err := deleteIfNecessary(users, name)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if deleted {
// 		fmt.Println("Deleted:", name)
// 		return
// 	}
// 	fmt.Println("Did not delete:", name)
// }

// func main() {
// 	users := map[string]user{
// 		"john": {
// 			name:                 "john",
// 			number:               18965554631,
// 			scheduledForDeletion: true,
// 		},
// 		"elon": {
// 			name:                 "elon",
// 			number:               19875556452,
// 			scheduledForDeletion: true,
// 		},
// 		"breanna": {
// 			name:                 "breanna",
// 			number:               98575554231,
// 			scheduledForDeletion: false,
// 		},
// 		"kade": {
// 			name:                 "kade",
// 			number:               10765557221,
// 			scheduledForDeletion: false,
// 		},
// 	}
// 	test(users, "john")
// 	test(users, "musk")
// 	test(users, "santa")
// 	test(users, "kade")

// 	keys := []string{}
// 	for name := range users {
// 		keys = append(keys, name)
// 	}
// 	sort.Strings(keys)

// 	fmt.Println("Final map keys:")
// 	for _, name := range keys {
// 		fmt.Println(" - ", name)
// 	}
// }

// ################################################################

// KEY TYPES
// Any type can be used as the value in a map, but keys are more restrictive.

// As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

// It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:

// hits := make(map[string]map[string]int)
// This is map of string to (map of string to int). Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:

// n := hits["/doc/"]["au"]
// Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:

// func add(m map[string]map[string]int, path, country string) {
//     mm, ok := m[path]
//     if !ok {
//         mm = make(map[string]int)
//         m[path] = mm
//     }
//     mm[country]++
// }
// add(hits, "/doc/", "au")
// On the other hand, a design that uses a single map with a struct key does away with all that complexity:

// type Key struct {
//     Path, Country string
// }
// hits := make(map[Key]int)
// When a Vietnamese person visits the home page, incrementing (and possibly creating) the appropriate counter is a one-liner:

// hits[Key{"/", "vn"}]++
// And it’s similarly straightforward to see how many Swiss people have read the spec:

// n := hits[Key{"/ref/spec", "ch"}]

// ===============================

func getCounts(userIDs []string) map[string]int {
	// ?
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

// don't edit below this line

func test(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}

func main() {
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	test(userIDs, []string{"00", "ff", "dd"})
	test(userIDs, []string{"aa", "12", "32"})
	test(userIDs, []string{"bb", "33"})
}
 