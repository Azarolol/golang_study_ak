package main

import "fmt"

func main() {
	var (
		name, city string
		age        int
	)
	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	fmt.Print("Введите ваш город: ")
	fmt.Scanln(&city)
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Город: %s\n", city)
}
