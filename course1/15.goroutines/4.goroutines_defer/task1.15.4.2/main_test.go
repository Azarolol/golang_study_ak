package main

import (
	"reflect"
	"testing"
)

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

func Test_testDefer(t *testing.T) {
	type args struct {
		a *int
		b *int
	}
	a := 8
	b := 13
	c := 104
	tests := []struct {
		name string
		args args
		want *int
	}{
		{"Test defer", args{&a, &b}, &c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testDefer(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDefer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test sum", args{8, 13}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test multiply", args{8, 13}, 104},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
