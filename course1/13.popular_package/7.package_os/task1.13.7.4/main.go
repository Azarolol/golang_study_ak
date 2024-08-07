package main

import (
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	output, err := exec.Command(binPath, args...).Output()
	if err != nil {
		return "Error executing binary: " + err.Error()
	}
	return string(output)
}

func main() {
	fmt.Println(ExecBin("C:/Users/Азамат/go/src/student.vkusvill.ru/Azarolol/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.4/test/test.exe"))
	fmt.Println(ExecBin("nonexistent-binary"))
}
