package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	dict := make(map[string]struct{})
	newUsers := make([]User, 0)
	for _, user := range users {
		_, ok := dict[user.Nickname]
		if !ok {
			newUsers = append(newUsers, user)
			dict[user.Nickname] = struct{}{}
		}
	}
	output := make([]User, len(newUsers))
	copy(output, newUsers)
	return output
}

func main() {
	users := []User{
		{"Gimazov", 23, "gimazovazamat@mail.ru"},
		{"Gimazov 2", 15, "gimazovalmaz@mail.ru"},
		{"Gimazova", 11, "gimazovaalisa@mail.ru"},
		{"Gimazov", 48, "gimazovramil@mail.ru"},
	}
	uniqueUsers := getUniqueUsers(users)
	fmt.Println(uniqueUsers, len(uniqueUsers), cap(uniqueUsers))
}
