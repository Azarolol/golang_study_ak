package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialFaster() map[string]bool {
	begin := time.Now()
	_ = factorialIterative(100000)
	iterativeTime := time.Since(begin)
	begin = time.Now()
	_ = factorialRecursive(100000)
	recursiveTime := time.Since(begin)
	out := make(map[string]bool)
	if iterativeTime < recursiveTime {
		out["Iterative"] = true
		out["Recursive"] = false
	} else if iterativeTime > recursiveTime {
		out["Iterative"] = false
		out["Recursive"] = true
	} else {
		out["Iterative"] = true
		out["Recursive"] = true
	}
	return out
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialFaster())
}
