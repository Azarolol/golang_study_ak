package main

import (
	"sync"
	"testing"

	"math/rand"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

var persons = make([]*Person, 0)

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = &Person{rand.Intn(100)}
		persons = append(persons, p)
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = &Person{rand.Intn(100)}
		personPool.Put(p)
	}
}

/*
BenchmarkWithoutPool-12    	15383983	        89.18 ns/op	      55 B/op	       1 allocs/op
BenchmarkWithPool-12    	22472667	        56.60 ns/op	      37 B/op	       1 allocs/op
*/
