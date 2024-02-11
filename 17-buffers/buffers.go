package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
Buffers

	Variable-sized buffer of bytes.
	Its zero value is an empty buffer.
*/
func main() {
	var b bytes.Buffer // buffer needs no initialization
	// write bytes to the buffer
	b.Write([]byte("Hello "))
	// write string literals to the buffer using WriteString()
	b.WriteString("world")
	// write using fmt.Fprintf(): formats & writes to buffer
	fmt.Fprintf(&b, "!\n")
	// get buffer size (string length)
	fmt.Println("buffer size:", b.Len())
	// get string value without draining
	fmt.Println("printed:", b.String())
	// drains buffer & write it to stdout (print)
	b.WriteTo(os.Stdout)
	// grows buffer's capacity, if necessary, to guarantee space for another n bytes
	b.Grow(64) // adds 64 extra bytes capacity
}
