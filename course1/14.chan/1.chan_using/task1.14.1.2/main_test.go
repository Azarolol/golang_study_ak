package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_generateData(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test generate data", args{10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateData(tt.args.n)
			time.Sleep(time.Second)
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("generateData() = %v, want %v", got, tt.want)
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
