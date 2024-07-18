package main

import "fmt"

func ConcatenateString(sep string, str ...string) string {
	even := "even: "
	odd := "odd: "
	for i, s := range str {
		if i%2 == 0 {
			if i > 0 {
				even += sep
			}
			even += s
		} else {
			if i > 1 {
				odd += sep
			}
			odd += s
		}
	}
	return even + ", " + odd
}

func main() {
	fmt.Print(ConcatenateString("-", "hello", "world", "how", "are", "you"))
}
