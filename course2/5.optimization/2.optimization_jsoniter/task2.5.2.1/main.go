package main

import (
	"encoding/json"
	"jsonBenchmark/user"

	"github.com/brianvoe/gofakeit/v6"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

func EasyJSON(users []user.User) {
	for _, user := range users {
		easyjson.Marshal(user)
	}
}

func JSON(users []user.User) {
	for _, user := range users {
		json.Marshal(user)
	}
}

func JSONiter(users []user.User) {
	for _, user := range users {
		jsoniter.Marshal(user)
	}
}

func GenerateUser(count int) []user.User {
	users := make([]user.User, count)
	for i := 0; i < count; i++ {
		users[i] = user.User{
			ID:       i,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 14),
			Age:      gofakeit.Number(18, 100),
			Email:    gofakeit.Email(),
		}
	}

	return users
}
