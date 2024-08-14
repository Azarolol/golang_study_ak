package main

import "testing"

func Test_generateMathString(t *testing.T) {
	type args struct {
		operands []int
		operator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test generate math string 1", args{[]int{2, 4, 6}, "+"}, "2 + 4 + 6 = 12"},
		{"Test generate math string 2", args{[]int{6, 4, 2}, "-"}, "6 - 4 - 2 = 0"},
		{"Test generate math string 3", args{[]int{2, 4, 6}, "*"}, "2 * 4 * 6 = 48"},
		{"Test generate math string 4", args{[]int{18, 6, 3}, "/"}, "18 / 6 / 3 = 1"},
		{"Test generate math string 5", args{[]int{}, "+"}, ""},
		{"Test generate math string 6", args{[]int{2, 4, 6}, "?"}, "Unknown operator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMathString(tt.args.operands, tt.args.operator); got != tt.want {
				t.Errorf("generateMathString() = %v, want %v", got, tt.want)
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
