package main

import "testing"

func TestMaxDifference(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test max difference 1", args{[]int{3, 4, 5, 0, 5, 3, 6, 2, 1}}, 6},
		{"Test max difference 2", args{[]int{}}, 0},
		{"Test max difference 3", args{[]int{5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDifference(tt.args.numbers); got != tt.want {
				t.Errorf("MaxDifference() = %v, want %v", got, tt.want)
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
