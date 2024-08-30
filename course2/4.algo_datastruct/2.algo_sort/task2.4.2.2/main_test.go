package main

import (
	"reflect"
	"testing"
)

func TestGeneralSort(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"Test general sort", []int{2, 11, 12, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []int{64, 34, 25, 12, 2, 11, 90}
			GeneralSort(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("generalSort() = %v, want %v", data, tt.want)
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

func Test_selectionSort(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"Test selection sort", []int{2, 11, 12, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []int{64, 34, 25, 12, 2, 11, 90}
			selectionSort(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", data, tt.want)
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"Test insertion sort", []int{2, 11, 12, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []int{64, 34, 25, 12, 2, 11, 90}
			insertionSort(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("inseriongSort() = %v, want %v", data, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"Test quick sort", []int{2, 11, 12, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []int{64, 34, 25, 12, 2, 11, 90}
			quickSort(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("quickSort() = %v, want %v", data, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test merge sort", args{[]int{64, 34, 25, 12, 2, 11, 90}}, []int{2, 11, 12, 25, 34, 64, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
