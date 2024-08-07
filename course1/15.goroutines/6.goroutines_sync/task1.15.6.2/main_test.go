package main

import "testing"

func TestCounter_Increment(t *testing.T) {
	tests := []struct {
		name string
		c    *Counter
		want int
	}{
		{"Test counter increment", &Counter{}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Increment(); got != tt.want {
				t.Errorf("Counter.Increment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concurrentSafeCounter(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test concurrent safe counter", 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concurrentSafeCounter(); got != tt.want {
				t.Errorf("concurrentSafeCounter() = %v, want %v", got, tt.want)
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
