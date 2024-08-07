package main

import (
	"reflect"
	"testing"
)

func TestRemoveIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test remove 1", args{[]int{0, 1, 2, 3, 4, 5}, 4}, []int{0, 1, 2, 3, 5}},
		{"Test remove 2", args{[]int{0, 1, 2, 3, 4, 5}, 6}, []int{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIDX(tt.args.xs, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIDX() = %v, want %v", got, tt.want)
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
