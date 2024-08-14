package main

import (
	"reflect"
	"testing"
)

func TestInsertToStart(t *testing.T) {
	type args struct {
		xs []int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test insert to start 1", args{[]int{1, 2, 3}, []int{4, 5, 6}}, []int{4, 5, 6, 1, 2, 3}},
		{"Test insert to start 2", args{[]int{1, 2, 3}, []int{}}, []int{1, 2, 3}},
		{"Test insert to start 3", args{[]int{}, []int{4, 5, 6}}, []int{4, 5, 6}},
		{"Test insert to start 4", args{[]int{}, []int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertToStart(tt.args.xs, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertToStart() = %v, want %v", got, tt.want)
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
