package main

import (
	"reflect"
	"testing"
)

func Test_countRussianLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{"Test count russian letters", args{"Привет, мир!"}, map[rune]int{
			'п': 1,
			'р': 2,
			'и': 2,
			'в': 1,
			'е': 1,
			'т': 1,
			'м': 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRussianLetters(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRussianLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRussianLetter(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test is russian letter 1", args{'g'}, false},
		{"Test is russian letter 2", args{'г'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRussianLetter(tt.args.r); got != tt.want {
				t.Errorf("isRussianLetter() = %v, want %v", got, tt.want)
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
