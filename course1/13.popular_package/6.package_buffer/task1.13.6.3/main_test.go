package main

import (
	"bytes"
	"testing"
)

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

func Test_getDataString(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test get data string", args{bytes.NewBufferString("Hello, World!")}, "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDataString(tt.args.buffer); got != tt.want {
				t.Errorf("getDataString() = %v, want %v", got, tt.want)
			}
		})
	}
}
