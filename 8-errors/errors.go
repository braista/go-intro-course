package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Error interface

	Go programs express errors with error values
	Is any type that implements the built-in error interface

	type error interface {
		Error() string
	}

	Code that calls a function that could return an error should handle it by testing if "error" is nil
*/
func main() {
	i, err := strconv.Atoi("42b")
	if err != nil { // checking it error is not nil
		fmt.Println("couldn't convert:", err)
		return
	}
	fmt.Println(i)

	// calling out custom function that can throw an error
	result, err := problematicFn(1)
	if err != nil {
		fmt.Println("it's OK! result:", result)
	}

	// checking custom error
	if err = sendSMS("hello", "user1234"); err != nil {
		fmt.Println(err)
	}
}

// function that can throw an error
func problematicFn(param1 int) (result int, err error) {
	if param1 < 10 {
		// NOTE: error strings shouldn't be capitalized!
		return 0, errors.New("number has to be greater than 10") // return error as second "return"
	}
	// everything's ok, so we return error as nil
	return param1 * 10, nil
}

// creating a custom error type
type userError struct {
	username string
}

func (e userError) Error() string { // implementing error interface
	return fmt.Sprintf("%v has a problem with their account", e.username)
}

func sendSMS(msg, userName string) error {
	canSend := true // check if can be sent
	if !canSend {
		return userError{username: userName} // creating an instance of our custom error
	}
	// send sms
	return nil
}
