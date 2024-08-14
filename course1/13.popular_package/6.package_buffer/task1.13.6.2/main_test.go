package main

import (
	"bytes"
	"reflect"
	"strings"
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

func Test_getScanner(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test get scanner 1", args{bytes.NewBuffer([]byte("Hello\n,\nWorld!"))}, "Hello\n,\nWorld!"},
		{"Test get scanner 2", args{bytes.NewBuffer([]byte{})}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ""
			scanner := getScanner(tt.args.b)
			for scanner.Scan() {
				output += scanner.Text() + "\n"
			}
			if got := strings.TrimRight(output, "\n"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
