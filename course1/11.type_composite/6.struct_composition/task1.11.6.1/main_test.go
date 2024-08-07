package main

import (
	"reflect"
	"testing"
)

func TestOrder_AddDish(t *testing.T) {
	type args struct {
		dish Dish
	}
	tests := []struct {
		name  string
		order *Order
		args  args
		want  *Order
	}{
		{"Test add dish", &Order{Dishes: []Dish{}, Total: 0.0}, args{Dish{"Pizza", 10.99}}, &Order{Dishes: []Dish{{"Pizza", 10.99}}, Total: 0.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.AddDish(tt.args.dish)
			if got := tt.order; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_RemoveDish(t *testing.T) {
	type args struct {
		dish Dish
	}
	tests := []struct {
		name  string
		order *Order
		args  args
		want  *Order
	}{
		{"Test remove dish", &Order{Dishes: []Dish{{"Pizza", 10.99}}, Total: 0.0}, args{Dish{"Pizza", 10.99}}, &Order{Dishes: []Dish{}, Total: 0.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.RemoveDish(tt.args.dish)
			if got := tt.order; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_CalculateTotal(t *testing.T) {
	tests := []struct {
		name  string
		order *Order
		want  float64
	}{
		{"Test calculate total", &Order{Dishes: []Dish{{"Pizza", 10.99}, {"Burger", 5.99}}}, 16.98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.CalculateTotal()
			if got := tt.order.Total; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.want)
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
