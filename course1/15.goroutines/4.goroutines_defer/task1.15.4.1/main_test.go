package main

import (
	"os"
	"testing"
)

func Test_writeToFile(t *testing.T) {
	type args struct {
		file *os.File
		data string
	}
	file, _ := os.OpenFile("file.txt", os.O_RDWR, 0777)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test write to file", args{file, "Hello, world!"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeToFile(tt.args.file, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeToFile() error = %v, wantErr %v", err, tt.wantErr)
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
