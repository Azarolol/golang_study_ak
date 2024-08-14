package main

import (
	"bytes"
	"fmt"
)

func getReader(b *bytes.Buffer) *bytes.Reader {
	return bytes.NewReader(b.Bytes())
}

func main() {
	// Create a buffer for testing
	buffer := bytes.NewBufferString("Hello, World!")
	b := make([]byte, 13)
	r := getReader(buffer)
	r.Read(b)
	fmt.Println(string(b)) // Hello, World!
}
