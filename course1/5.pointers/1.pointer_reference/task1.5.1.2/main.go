package main

import "fmt"

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	newStr := ""
	for _, c := range *str {
		newStr = string(c) + newStr
	}
	*str = newStr
}

func main() {
	a := 41
	mutate(&a)
	fmt.Println(a)
	str := "string to be reversed"
	ReverseString(&str)
	fmt.Println(str)
}
