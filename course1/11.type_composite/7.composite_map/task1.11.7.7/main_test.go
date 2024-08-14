package main

import (
	"reflect"
	"testing"
)

func Test_getUniqueUsers(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"Test get unique users", args{[]User{
			{"Gimazov", 23, "gimazovazamat@mail.ru"},
			{"Gimazov 2", 15, "gimazovalmaz@mail.ru"},
			{"Gimazova", 11, "gimazovaalisa@mail.ru"},
			{"Gimazov", 48, "gimazovramil@mail.ru"},
		}}, []User{
			{"Gimazov", 23, "gimazovazamat@mail.ru"},
			{"Gimazov 2", 15, "gimazovalmaz@mail.ru"},
			{"Gimazova", 11, "gimazovaalisa@mail.ru"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUniqueUsers() = %v, want %v", got, tt.want)
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
