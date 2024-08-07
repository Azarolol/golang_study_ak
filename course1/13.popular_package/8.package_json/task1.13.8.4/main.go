package main

import (
	"encoding/json"
	"os"
	"strings"
)

func writeJSON(filePath string, data interface{}) error {
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
	err := writeJSON("course1/13.popular_package/8.package_json/task1.13.8.4/test/users.json", data)
	if err != nil {
		panic(err)
	}
}
