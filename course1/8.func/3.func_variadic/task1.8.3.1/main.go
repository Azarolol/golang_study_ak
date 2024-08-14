package main

import "fmt"

func UserInfo(name string, age int, cities ...string) string {
	citiesString := ""
	for i, city := range cities {
		if i > 0 {
			citiesString += ", "
		}
		citiesString += city
	}
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %s", name, age, citiesString)
}

func main() {
	fmt.Println(UserInfo("John", 21, "Moscow", "Saint Petersburg"))
	fmt.Println(UserInfo("Alex", 34))
	fmt.Println(UserInfo("Иван", 25, "Москва", "Санкт-Петербург", "Казань"))
}
