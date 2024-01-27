package main

import "fmt"

/*
	Maps
		Data structure that associate values of one type (key) with values of another type (value).
		The key can be of any type for which the equality operator is defined.

		Zero value: nil
*/
func main() {
	// maps can be constructed using composite literal syntax
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Println(timeZone)

	// or can be constructed using "make" function
	var studentNotes = make(map[string]int)

	// add a new key-value pair
	studentNotes["Peter"] = 8
	studentNotes["Richard"] = 3
	fmt.Println(studentNotes) // prints map[Peter:8]

	// get an element (value) by key
	fmt.Println(studentNotes["Peter"])
	fmt.Println(studentNotes["Carlos"]) // doesn't exist.. but CAREFUL, still returns zero value

	// delete an element by key
	delete(studentNotes, "Richard")

	// check if key exists (use to differenciate from zero value)
	_, ok := studentNotes["Richard"]
	fmt.Println("Does element exist?", ok)

	// checking map size
	fmt.Println(len(studentNotes))
}
