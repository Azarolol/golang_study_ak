package main

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"Test pop 1", args{[]int{1, 2, 3, 4, 5}}, 1, []int{2, 3, 4, 5}},
		{"Test pop 2", args{[]int{}}, 0, []int{}},
		{"Test pop 3", args{[]int{1}}, 1, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pop(tt.args.xs)
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
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
