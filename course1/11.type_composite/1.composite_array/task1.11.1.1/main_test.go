package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		xs [8]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test sum 1", args{[8]int{1, 2, 3, 4, 5, 6, 7, 8}}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.xs); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_average(t *testing.T) {
	type args struct {
		xs [8]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test average", args{[8]int{1, 2, 3, 4, 5, 6, 7, 8}}, 4.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.xs); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageFloat(t *testing.T) {
	type args struct {
		ys [8]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test average float", args{[8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageFloat(tt.args.ys); got != tt.want {
				t.Errorf("averageFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		xs [8]int
	}
	tests := []struct {
		name string
		args args
		want [8]int
	}{
		{"Test reverse", args{[8]int{1, 2, 3, 4, 5, 6, 7, 8}}, [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Main test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
