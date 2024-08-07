package main

import (
	"reflect"
	"testing"
)

func Test_countWordOccurences(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"Test count word occurences", args{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"}, map[string]int{"Lorem": 1, "ipsum": 2, "dolor": 1,
			"sit": 1, "amet": 1, "consectetur": 1, "adipiscing": 1, "elit": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWordOccurences(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWordOccurences() = %v, want %v", got, tt.want)
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
