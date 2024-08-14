package main

import (
	"encoding/json"
	"fmt"
)

func getJSON(data []User) (string, error) {
	out, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func main() {
	users := []User{
		{"Азамат", 23, []Comment{
			{"Первый комментарий"},
			{"Второй комментарий"},
		}},
		{"Алмаз", 15, []Comment{
			{"Комментарий"},
		},
		},
	}
	output, err := getJSON(users)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
