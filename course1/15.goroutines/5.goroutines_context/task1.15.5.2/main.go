package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithTimeout(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithTimeout(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
}

func contextWithTimeout(ctx context.Context, contextTimeout time.Duration, timeAfter time.Duration) string {
	ctxWT, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	for {
		select {
		case <-ctxWT.Done():
			return "превышено время ожидания контекста"
		case <-time.After(timeAfter):
			return "превышено время ожидания"
		}
	}
}
