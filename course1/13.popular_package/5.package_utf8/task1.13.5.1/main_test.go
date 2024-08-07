package main

import "testing"

func Test_countUniqueUTF8Chars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test count unique UT8 chars 1", args{"Hel1o,	 !"}, 9},
		{"Test count unique UT8 chars 2", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueUTF8Chars(tt.args.s); got != tt.want {
				t.Errorf("countUniqueUTF8Chars() = %v, want %v", got, tt.want)
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
