package main

import (
	"testing"
)

func BenchmarkEasyJSON(b *testing.B) {
	users := GenerateUser(100000)
	b.ResetTimer()
	EasyJSON(users)
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkEasyJSON$ jsonBenchmark

goos: windows
goarch: amd64
pkg: jsonBenchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkEasyJSON-12    	1000000000	         0.04788 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	jsonBenchmark	3.019s
*/

func BenchmarkJSON(b *testing.B) {
	users := GenerateUser(100000)
	b.ResetTimer()
	JSON(users)
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkJSON$ jsonBenchmark

goos: windows
goarch: amd64
pkg: jsonBenchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkJSON-12    	1000000000	         0.1373 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	jsonBenchmark	5.290s
*/

func BenchmarkJSONiter(b *testing.B) {
	users := GenerateUser(100000)
	b.ResetTimer()
	JSONiter(users)
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkJSONiter$ jsonBenchmark

goos: windows
goarch: amd64
pkg: jsonBenchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkJSONiter-12    	1000000000	         0.08429 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	jsonBenchmark	3.100s
*/
