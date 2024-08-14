package main

import "testing"

func Test_bitwiseXOR(t *testing.T) {
	type args struct {
		n   int
		res int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test bitwise XOR 1", args{6, 6}, 0},
		{"Test bitwise XOR 2", args{16, 8}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseXOR(tt.args.n, tt.args.res); got != tt.want {
				t.Errorf("bitwiseXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSingleNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test find single number 1", args{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}}, 5},
		{"Test find single number 2", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSingleNumber(tt.args.numbers); got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
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
