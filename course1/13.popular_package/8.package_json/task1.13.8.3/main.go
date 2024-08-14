package main

import (
	"encoding/json"
	"os"
	"strings"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func writeJSON(filePath string, data []User) error {
	if strings.Contains(filePath, "/") {
		dirs := strings.Split(filePath, "/")
		err := os.MkdirAll(strings.TrimSuffix(filePath, "/"+dirs[len(dirs)-1]), 0777)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(filePath, os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func main() {
	err := writeJSON("course1/13.popular_package/8.package_json/task1.13.8.3/test/file.json", []User{
		{"Азамат", 23, []Comment{
			{"Первый комментарий"},
			{"Второй комментарий"},
		}},
		{"Алмаз", 15, []Comment{
			{"Комментарий"},
		},
		},
	})
	if err != nil {
		panic(err)
	}
}
