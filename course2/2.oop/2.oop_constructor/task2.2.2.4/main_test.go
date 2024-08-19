package main

import (
	"reflect"
	"testing"
)

func TestNewLogSystem(t *testing.T) {
	type args struct {
		lops []LogOption
	}
	tests := []struct {
		name string
		args args
		want *LogSystem
	}{
		{"Test new log system", args{[]LogOption{}}, &LogSystem{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogSystem(tt.args.lops...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogSystem() = %v, want %v", got, tt.want)
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
