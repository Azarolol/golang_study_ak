package main

import "testing"

func TestReplaceSymbols(t *testing.T) {
	type args struct {
		str string
		old rune
		new rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test replace symbols", args{"Hello, world!", 'o', '0'}, "Hell0, w0rld!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSymbols(tt.args.str, tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("ReplaceSymbols() = %v, want %v", got, tt.want)
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
