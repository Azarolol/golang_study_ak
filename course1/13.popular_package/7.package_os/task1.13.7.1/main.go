package main

import (
	"os"
	"strings"
)

func WriteFile(filepath string, data []byte, perm os.FileMode) error {
	if strings.Contains(filepath, "/") {
		dirs := strings.Split(filepath, "/")
		err := os.MkdirAll(strings.TrimSuffix(filepath, "/"+dirs[len(dirs)-1]), perm)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(filepath, data, perm)
}

func main() {
	err := WriteFile("course1/13.popular_package/7.package_os/task1.13.7.1/test/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
