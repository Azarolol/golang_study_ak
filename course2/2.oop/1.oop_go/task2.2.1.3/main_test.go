package main

import (
	"reflect"
	"testing"
)

func TestDineInOrder_AddItem(t *testing.T) {
	type args struct {
		item     string
		quantity int
	}
	tests := []struct {
		name    string
		o       *DineInOrder
		args    args
		wantErr bool
	}{
		{"Test add item to dine in order", &DineInOrder{make(map[string]int)}, args{"Pizza", 1}, false},
		{"Test add item to dine in order (quantity less than 1)", &DineInOrder{make(map[string]int)}, args{"Pizza", 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.AddItem(tt.args.item, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("DineInOrder.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTakeAwayOrder_AddItem(t *testing.T) {
	type args struct {
		item     string
		quantity int
	}
	tests := []struct {
		name    string
		o       *TakeAwayOrder
		args    args
		wantErr bool
	}{
		{"Test add item to take away order", &TakeAwayOrder{make(map[string]int)}, args{"Pizza", 1}, false},
		{"Test add item to take away order (quantity less than 1)", &TakeAwayOrder{make(map[string]int)}, args{"Pizza", 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.AddItem(tt.args.item, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("TakeAwayOrder.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDineInOrder_RemoveItem(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name    string
		o       *DineInOrder
		args    args
		wantErr bool
	}{
		{"Test remove item from dine in order", &DineInOrder{map[string]int{"Pizza": 1}}, args{"Pizza"}, false},
		{"Test remove item from dine in order (no such item in order)", &DineInOrder{map[string]int{"Pizza": 1}}, args{"Burger"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.RemoveItem(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("DineInOrder.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTakeAwayOrder_RemoveItem(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name    string
		o       *TakeAwayOrder
		args    args
		wantErr bool
	}{
		{"Test remove item from take away order", &TakeAwayOrder{map[string]int{"Pizza": 1}}, args{"Pizza"}, false},
		{"Test remove item from take away order (no such item in order)", &TakeAwayOrder{map[string]int{"Pizza": 1}}, args{"Burger"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.RemoveItem(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("TakeAwayOrder.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDineInOrder_GetOrderDetails(t *testing.T) {
	tests := []struct {
		name string
		o    *DineInOrder
		want map[string]int
	}{
		{"Test get dine in order details", &DineInOrder{map[string]int{"Pizza": 1}}, map[string]int{"Pizza": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOrderDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DineInOrder.GetOrderDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTakeAwayOrder_GetOrderDetails(t *testing.T) {
	tests := []struct {
		name string
		o    *TakeAwayOrder
		want map[string]int
	}{
		{"Test take away order details", &TakeAwayOrder{map[string]int{"Pizza": 1}}, map[string]int{"Pizza": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOrderDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TakeAwayOrder.GetOrderDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManageOrder(t *testing.T) {
	type args struct {
		o Order
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test manage dine in order", args{&DineInOrder{make(map[string]int)}}},
		{"Test manage take away order", args{&TakeAwayOrder{make(map[string]int)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ManageOrder(tt.args.o)
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
