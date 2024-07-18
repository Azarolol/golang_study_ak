package main

import "fmt"

var (
	name, surname, age, city string
	ageInt, year, month, day int
)

func main() {
	name = "Azamat"
	surname = "Gimazov"
	age = "23"
	city = "Moscow"
	ageInt = 23
	year = 2001
	month = 3
	day = 29
	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(age)
	fmt.Println(city)
	fmt.Println(ageInt)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
}
