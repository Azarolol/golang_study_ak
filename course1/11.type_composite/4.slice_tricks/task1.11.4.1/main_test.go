package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	type args struct {
		xs    []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test cut 1", args{[]int{1, 2, 3, 4, 5}, 1, 3}, []int{2, 3, 4}},
		{"Test cut 2", args{[]int{}, 1, 3}, []int{}},
		{"Test cut 3", args{[]int{}, -1, 3}, []int{}},
		{"Test cut 4", args{[]int{}, 1, 8}, []int{}},
		{"Test cut 5", args{[]int{}, 3, 1}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
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
