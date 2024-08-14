package main

import "testing"

func Test_trySend(t *testing.T) {
	type args struct {
		ch chan int
		v  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test try send 1", args{make(chan int, 1), 1}, true},
		{"Test try send 2", args{make(chan int), 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trySend(tt.args.ch, tt.args.v); got != tt.want {
				t.Errorf("trySend() = %v, want %v", got, tt.want)
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
