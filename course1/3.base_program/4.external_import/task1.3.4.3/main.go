package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	name := fake.FullName()
	address := fake.StreetAddress()
	phone := fake.Phone()
	email := fake.EmailAddress()
	return "Name: " + name + "\n" + "Address: " + address + "\n" + "Phone: " + phone + "\n" + "Email: " + email
}

func main() {
	fmt.Print(GenerateFakeData())
}
