package main

import "testing"

func TestExecBin(t *testing.T) {
	type args struct {
		binPath string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test exec bin 1", args{"C:/Users/Азамат/go/src/student.vkusvill.ru/Azarolol/go-course/golang_study_ak/course1/13.popular_package/7.package_os/task1.13.7.4/test/test.exe", []string{}}, "It is working!"},
		{"Test exec bin 2", args{"nonexistent-binary", []string{}}, "Error executing binary: exec: \"nonexistent-binary\": executable file not found in %PATH%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecBin(tt.args.binPath, tt.args.args...); got != tt.want {
				t.Errorf("ExecBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
