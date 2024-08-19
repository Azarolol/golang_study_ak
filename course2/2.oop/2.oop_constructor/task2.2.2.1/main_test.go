package main

import (
	"reflect"
	"testing"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		id   int
		oops []OrderOption
	}
	tests := []struct {
		name string
		args args
		want *Order
	}{
		{"Test new order", args{1, []OrderOption{}}, &Order{ID: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(tt.args.id, tt.args.oops...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
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
