package main

import (
	"reflect"
	"testing"
)

func Test_getUsersFromJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}{
		{"Test get users from JSON", args{[]byte(`[
		{
			"name":"John",
			"age":30,
			"comments":[
				{"text":"Great post!"},
				{"text":"I agree"}
			]
		},
		{
			"name":"Alice",
			"age":25,
			"comments":[
				{"text":"Nice article"}
			]
		}
	]`)}, []User{
			{"John", 30, []Comment{{"Great post!"}, {"I agree"}}},
			{"Alice", 25, []Comment{{"Nice article"}}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUsersFromJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUsersFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUsersFromJSON() = %v, want %v", got, tt.want)
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
