package main

import "fmt"

func concatStrings(xs ...string) string {
	output := ""
	for _, x := range xs {
		output += x
	}
	return output
}

func main() {
	result := concatStrings("Hello", " ", "world!")
	fmt.Println(result) // Вывод: "Hello world!"
}
