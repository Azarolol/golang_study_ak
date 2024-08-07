package main

import "testing"

func Test_concatStrings(t *testing.T) {
	type args struct {
		xs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test concat strings", args{[]string{"Hello", " ", "world!"}}, "Hello world!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatStrings(tt.args.xs...); got != tt.want {
				t.Errorf("concatStrings() = %v, want %v", got, tt.want)
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
