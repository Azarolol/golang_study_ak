package main

import (
	"reflect"
	"testing"
)

func Test_sortDescInt(t *testing.T) {
	type args struct {
		xs [8]int
	}
	tests := []struct {
		name string
		args args
		want [8]int
	}{
		{"Test sort int desc", args{[8]int{5, 2, 8, 1, 9, 3, 7, 4}}, [8]int{9, 8, 7, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescInt(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortAscInt(t *testing.T) {
	type args struct {
		xs [8]int
	}
	tests := []struct {
		name string
		args args
		want [8]int
	}{
		{"Test sort int asc", args{[8]int{5, 2, 8, 1, 9, 3, 7, 4}}, [8]int{1, 2, 3, 4, 5, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscInt(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDescFloat(t *testing.T) {
	type args struct {
		xs [8]float64
	}
	tests := []struct {
		name string
		args args
		want [8]float64
	}{
		{"Test sort float desc", args{[8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}}, [8]float64{9.9, 8.8, 7.7, 5.5, 4.4, 3.3, 2.2, 1.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescFloat(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortAscFloat(t *testing.T) {
	type args struct {
		xs [8]float64
	}
	tests := []struct {
		name string
		args args
		want [8]float64
	}{
		{"Test sort float asc", args{[8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}}, [8]float64{1.1, 2.2, 3.3, 4.4, 5.5, 7.7, 8.8, 9.9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscFloat(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscFloat() = %v, want %v", got, tt.want)
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
