package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		users[i] = User{gofakeit.Name(), gofakeit.Number(18, 60)}
	}
	return users
}

func preparePrint(users []User) string {
	output := ""
	for _, user := range users {
		output += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return output
}

func main() {
	// Пример использования функции
	users := getUsers() // Получаем срез 10 пользователей
	result := preparePrint(users)
	fmt.Println(result)
}
