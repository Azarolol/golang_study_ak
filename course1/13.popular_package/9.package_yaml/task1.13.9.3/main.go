package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data []User) error {
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
	encoder := yaml.NewEncoder(file)
	return encoder.Encode(data)
}

func main() {
	err := writeYAML("course1/13.popular_package/9.package_yaml/task1.13.9.3/test/file.yaml", []User{
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
