package main

import (
	"reflect"
	"testing"
)

func TestRemoveExtraMemory(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test remove extra memory", args{make([]int, 3, 10)}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cap(RemoveExtraMemory(tt.args.xs)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveExtraMemory() = %v, want %v", got, tt.want)
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
