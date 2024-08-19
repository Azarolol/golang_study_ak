package main

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ID   int
		uops []UserOption
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{"Test new user", args{1, []UserOption{}}, &User{ID: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.ID, tt.args.uops...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
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
