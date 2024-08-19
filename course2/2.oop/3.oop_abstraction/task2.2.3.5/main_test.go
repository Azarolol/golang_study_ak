package main

import (
	"strconv"
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

func Benchmark64(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC64())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func Benchmark32(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC32())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func Benchmark16(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC16())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func Benchmark8(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC8())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}
