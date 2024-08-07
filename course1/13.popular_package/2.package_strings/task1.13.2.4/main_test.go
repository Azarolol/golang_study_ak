package main

import (
	"testing"
	"unicode/utf8"
)

func Test_generateActivationKey(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test generate activation key", 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utf8.RuneCountInString(generateActivationKey()); got != tt.want {
				t.Errorf("generateActivationKey() = %v, want %v", got, tt.want)
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
