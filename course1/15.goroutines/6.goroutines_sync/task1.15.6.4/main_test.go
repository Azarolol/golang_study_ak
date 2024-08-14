package main

import (
	"testing"
)

func Test_keyBuilder(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test key builder", args{[]string{"user", "1"}}, "user:1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyBuilder(tt.args.keys...); got != tt.want {
				t.Errorf("keyBuilder() = %v, want %v", got, tt.want)
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
