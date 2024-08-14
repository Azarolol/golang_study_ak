package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateSlice(size int) []float64 {
	output := make([]float64, size)
	for i := 0; i < size; i++ {
		output[i] = rand.Float64() * 100
	}
	return output
}

func TestAverage(t *testing.T) {
	input := generateSlice(10)
	test1 := average(input)
	test2 := average(input)
	result := reflect.DeepEqual(test1, test2)
	if !result {
		t.Errorf("Average test failed")
	}
}
