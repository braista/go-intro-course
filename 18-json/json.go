package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
JSON

	Simple data interchange format.
*/
func main() {
	type message struct {
		Title string    `json:"title"` // use this to set json field name
		Body  string    `json:"body"`
		Time  time.Time `json:"time"`
	}
	// marshalling
	// encoding JSON data using json.Marshal()
	m := message{"Alice", "Hello", time.Now()}
	body, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println("body:", string(body))

	// unmarshalling
	// decoding JSON data using json.Unmarshall()
	inputJSON := `{
		"title": "Some title",
		"body": "Some large text body",
		"time": "2024-02-11T17:47:53.395191-03:00"
	}`
	var decodedData message
	if err := json.Unmarshal([]byte(inputJSON), &decodedData); err != nil {
		panic(err)
	}
	fmt.Println("decoded data:", decodedData)
}
