package main

import (
	"reflect"
	"testing"
)

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

func TestFilterComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"Test filter comments", args{[]User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "BaD CoMmEnT 2"},
					{Message: "Bad Comment 3"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
					{Message: "Bad Comments 4"},
				},
			},
		}}, []User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBadComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test is bad comment 1", args{"good Comment 1"}, false},
		{"Test is bad comment 2", args{"BaD CoMmEnT 2"}, true},
		{"Test is bad comment 3", args{"Bad Comment 3"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadComment(tt.args.comment); got != tt.want {
				t.Errorf("IsBadComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBadComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{"Test get bad comments", args{[]User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "BaD CoMmEnT 2"},
					{Message: "Bad Comment 3"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
					{Message: "Bad Comments 4"},
				},
			},
		}}, []Comment{{Message: "BaD CoMmEnT 2"},
			{Message: "Bad Comment 3"},
			{Message: "Bad Comments 4"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBadComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoodComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{"Test get good comments", args{[]User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "BaD CoMmEnT 2"},
					{Message: "Bad Comment 3"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
					{Message: "Bad Comments 4"},
				},
			},
		}}, []Comment{
			{Message: "good Comment 1"},
			{Message: "Use camelCase please"},
			{Message: "Good Comment 1"},
			{Message: "Good Comment 2"},
			{Message: "Good Comment 3"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoodComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoodComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
