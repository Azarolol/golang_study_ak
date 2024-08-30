package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"Test merge", args{[]User{{1, "Azamat", 23}, {3, "Ramil", 49}, {5, "Irina", 46}}, []User{{2, "Almaz", 15}, {4, "Alisa", 11}}}, []User{{1, "Azamat", 23}, {2, "Almaz", 15}, {3, "Ramil", 49}, {4, "Alisa", 11}, {5, "Irina", 46}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
