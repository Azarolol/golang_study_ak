package main

import (
	"context"
	"testing"
	"time"
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

func Test_contextWithTimeout(t *testing.T) {
	type args struct {
		ctx            context.Context
		contextTimeout time.Duration
		timeAfter      time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test context with timeout 1", args{context.Background(), 1 * time.Second, 2 * time.Second}, "превышено время ожидания контекста"},
		{"Test context with timeout 2", args{context.Background(), 2 * time.Second, 1 * time.Second}, "превышено время ожидания"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contextWithTimeout(tt.args.ctx, tt.args.contextTimeout, tt.args.timeAfter); got != tt.want {
				t.Errorf("contextWithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
