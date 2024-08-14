package main

import (
	"reflect"
	"testing"
)

func Test_getUsers(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test get users", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(getUsers()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preparePrint(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test prepare print", args{[]User{{"Гимазов Азамат", 23}, {"Гимазов Рамиль", 48}}}, "Имя: Гимазов Азамат, Возраст: 23\nИмя: Гимазов Рамиль, Возраст: 48\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePrint(tt.args.users); got != tt.want {
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
