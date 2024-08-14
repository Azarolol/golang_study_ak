package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_getReader(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test get reader", args{bytes.NewBufferString("Hello, World!")}, "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 13)
			r := getReader(tt.args.b)
			r.Read(b)
			if got := string(b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReader() = %v, want %v", got, tt.want)
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
