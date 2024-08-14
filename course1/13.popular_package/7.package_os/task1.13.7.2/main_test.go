package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantFd  string
		wantErr bool
	}{
		{"Test write file", args{strings.NewReader("Hello, World!")}, "Hello, World!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := &bytes.Buffer{}
			if err := WriteFile(tt.args.data, fd); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFd := fd.String(); gotFd != tt.wantFd {
				t.Errorf("WriteFile() = %v, want %v", gotFd, tt.wantFd)
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
