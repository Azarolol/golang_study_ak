package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "test@example.com"
	valid := isValidEmail(email)
	if valid {
		fmt.Printf("%s является валидным email-адресом\n", email)
	} else {
		fmt.Printf("%s не явлется валидным email-адресом\n", email)
	}
}

func isValidEmail(email string) bool {
	output, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", email)
	return output
}
