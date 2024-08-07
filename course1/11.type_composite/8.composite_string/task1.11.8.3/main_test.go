package main

import (
	"reflect"
	"testing"
)

func Test_getBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Test get bytes", args{"Привет, мир!"}, []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 184, 209, 128, 33}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test get runes", args{"Привет, мир!"}, []rune{1055, 1088, 1080, 1074, 1077, 1090, 44, 32, 1084, 1080, 1088, 33}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRunes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRunes() = %v, want %v", got, tt.want)
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
