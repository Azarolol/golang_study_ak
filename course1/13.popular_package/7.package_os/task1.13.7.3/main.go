package main

import (
	"fmt"
	"os"
)

func ReadString(filePath string) string {
	output, err := os.ReadFile(filePath)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

func main() {
	fmt.Println(ReadString("course1/13.popular_package/7.package_os/task1.13.7.3/file.txt"))
}
