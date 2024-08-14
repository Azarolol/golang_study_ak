package main

import "testing"

func TestCountVowels(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test count vowels 1", args{"Привет, мир!"}, 3},
		{"Test count vowels 2", args{"Hello, world!"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountVowels(tt.args.str); got != tt.want {
				t.Errorf("CountVowels() = %v, want %v", got, tt.want)
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
