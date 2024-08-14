package main

import (
	"reflect"
	"testing"
)

func TestOperate(t *testing.T) {
	type args struct {
		f func(xs ...interface{}) interface{}
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Test operate concat", args{Concat, []interface{}{"Hello, ", "World!"}}, "Hello, World!"},
		{"Test operate sum int", args{Sum, []interface{}{1, 2, 3, 4, 5}}, 15},
		{"Test operate sum float", args{Sum, []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}}, 16.5},
		{"Test operate sum empty", args{Sum, []interface{}{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Operate(tt.args.f, tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Test concat 1", args{[]interface{}{"Hello, ", "World!"}}, "Hello, World!"},
		{"Test concat 2", args{[]interface{}{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Test sum int", args{[]interface{}{1, 2, 3, 4, 5}}, 15},
		{"Test sum float", args{[]interface{}{1.1, 2.2, 3.3, 4.4, 5.5}}, 16.5},
		{"Test sum empty", args{[]interface{}{}}, nil},
		{"Test sum nil", args{[]interface{}{"Hello, ", "World!"}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
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
