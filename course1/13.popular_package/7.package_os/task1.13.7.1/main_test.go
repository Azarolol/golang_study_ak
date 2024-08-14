package main

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		filepath string
		data     []byte
		perm     os.FileMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test write file", args{"test/file.txt", []byte("Hello, World!"), os.FileMode(0644)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.filepath, tt.args.data, tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
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
