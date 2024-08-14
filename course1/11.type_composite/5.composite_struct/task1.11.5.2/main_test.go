package main

import (
	"reflect"
	"testing"
)

func Test_getAnimals(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test get animals", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(getAnimals()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAnimals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preparePrint(t *testing.T) {
	type args struct {
		animals []Animal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test prepare print", args{[]Animal{{"Собака", "Жучка", 3}, {"Кошка", "Мята", 5}}}, "Тип: Собака, Имя: Жучка, Возраст: 3\nТип: Кошка, Имя: Мята, Возраст: 5\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePrint(tt.args.animals); got != tt.want {
				t.Errorf("preparePrint() = %v, want %v", got, tt.want)
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
