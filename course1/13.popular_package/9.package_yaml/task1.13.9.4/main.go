package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func writeYAML(filePath string, data interface{}) error {
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
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}
	err := writeYAML("course1/13.popular_package/9.package_yaml/task1.13.9.4/test/users.yaml", data)
	if err != nil {
		panic(err)
	}
}
