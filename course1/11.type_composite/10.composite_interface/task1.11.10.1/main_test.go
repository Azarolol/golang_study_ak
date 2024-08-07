package main

import "testing"

func Test_getType(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test get type 1", args{42}, "int"},
		{"Test get type 2", args{"Hello, World!"}, "string"},
		{"Test get type 3", args{[]int{1, 2, 3}}, "[]int"},
		{"Test get type 4", args{nil}, "Пустой интерфейс"},
		{"Test get type 5", args{42.2}, "Unknown type"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.i); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
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
