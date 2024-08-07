package main

import (
	"reflect"
	"testing"
)

func Test_getSubSlice(t *testing.T) {
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
		{"Test sub slice", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 6}, []int{3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubSlice(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubSlice() = %v, want %v", got, tt.want)
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
