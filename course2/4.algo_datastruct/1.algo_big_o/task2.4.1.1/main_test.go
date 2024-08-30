package main

import (
	"reflect"
	"testing"
)

func Test_factorialRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test factorial recursive 0", args{0}, 1},
		{"Test factorial recursive 1", args{1}, 1},
		{"Test factorial recursive 2", args{2}, 2},
		{"Test factorial recursive 3", args{3}, 6},
		{"Test factorial recursive 4", args{4}, 24},
		{"Test factorial recursive 5", args{5}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialRecursive(tt.args.n); got != tt.want {
				t.Errorf("factorialRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialIterative(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test factorial iterative 0", args{0}, 1},
		{"Test factorial iterative 1", args{1}, 1},
		{"Test factorial iterative 2", args{2}, 2},
		{"Test factorial iterative 3", args{3}, 6},
		{"Test factorial iterative 4", args{4}, 24},
		{"Test factorial iterative 5", args{5}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialIterative(tt.args.n); got != tt.want {
				t.Errorf("factorialIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareWhichFactorialFaster(t *testing.T) {
	tests := []struct {
		name string
		want map[string]bool
	}{
		{"Test compare which factorial faster", map[string]bool{"Iterative": true, "Recursive": false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareWhichFactorialFaster(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareWhichFactorialFaster() = %v, want %v", got, tt.want)
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
