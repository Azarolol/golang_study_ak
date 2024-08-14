package main

import "testing"

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

func Test_getVariableType(t *testing.T) {
	type args struct {
		variable interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test get variable type 1", args{10}, "int"},
		{"Test get variable type 2", args{"Hello"}, "string"},
		{"Test get variable type 3", args{interface{}(nil)}, "nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVariableType(tt.args.variable); got != tt.want {
				t.Errorf("getVariableType() = %v, want %v", got, tt.want)
			}
		})
	}
}
