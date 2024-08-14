package main

import (
	"reflect"
	"testing"
)

func TestInsertAfterIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
		x   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test insert after index 1", args{[]int{1, 2, 3, 4, 5}, 2, []int{6, 7, 8}}, []int{1, 2, 3, 6, 7, 8, 4, 5}},
		{"Test insert after index 2", args{[]int{1, 2, 3, 4, 5}, 5, []int{6, 7, 8}}, []int{}},
		{"Test insert after index 3", args{[]int{}, 0, []int{6, 7, 8}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAfterIDX(tt.args.xs, tt.args.idx, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAfterIDX() = %v, want %v", got, tt.want)
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
