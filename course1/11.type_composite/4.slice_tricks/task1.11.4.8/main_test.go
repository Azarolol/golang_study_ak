package main

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"Test shift 1", args{xs: []int{1, 2, 3, 4, 5}}, 1, []int{5, 1, 2, 3, 4}},
		{"Test shift 2", args{xs: []int{}}, 0, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Shift(tt.args.xs)
			if got != tt.want {
				t.Errorf("Shift() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Shift() got1 = %v, want %v", got1, tt.want1)
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
