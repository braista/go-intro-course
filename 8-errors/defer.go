package main

import (
	"fmt"
	"os"
)

/*
Defer use case

	Suppose we wanted to create a file, write to it, and then close when we’re done.
	Here’s how we could do that with defer

	https://gobyexample.com/defer
*/
func deferExample() {
	f := createFile("/tmp/defer.txt")
	/*
		Immediately after createFile, we defer the closing of that file with closeFile.
		This will be executed at the end of the enclosing function (main), after writeFile has finished
	*/
	defer closeFile(f) // defer closeFile to be called at the end, so it don't stays open
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
