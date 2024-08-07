package main

import (
	"os"
)

func writeToFile(file *os.File, data string) error {
	defer file.Close()
	_, err := file.WriteString(data)
	return err
}

func main() {
	data := "Hello, world!"
	file, err := os.OpenFile("file.txt", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	err = writeToFile(file, data)
	if err != nil {
		panic(err)
	}
}
