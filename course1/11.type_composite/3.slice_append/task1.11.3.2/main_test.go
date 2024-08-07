package main

import (
	"reflect"
	"testing"
)

func Test_appendInt(t *testing.T) {
	type args struct {
		xs *[]int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		{"Test append int", args{&[]int{0, 1, 2}, []int{3, 4, 5}}, &[]int{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		xs := tt.args.xs
		t.Run(tt.name, func(t *testing.T) {
			appendInt(xs, tt.args.x...)
			if got := xs; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendInt() = %v, want %v", got, tt.want)
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
