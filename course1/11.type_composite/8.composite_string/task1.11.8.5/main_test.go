package main

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test reverse string 1", args{"Hello, world!"}, "!dlrow ,olleH"},
		{"Test reverse string 2", args{"!"}, "!"},
		{"Test reverse string 3", args{""}, ""},
		{"Test reverse string 4", args{"   "}, "   "},
		{"Test reverse string 5", args{"\n?,"}, ",?\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.str); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
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
