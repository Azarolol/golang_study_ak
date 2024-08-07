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

func Test_contextWithDeadline(t *testing.T) {
	type args struct {
		ctx             context.Context
		contextDeadline time.Duration
		timeAfter       time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test context with deadline 1", args{context.Background(), 1 * time.Second, 2 * time.Second}, "context deadline exceeded"},
		{"Test context with deadline 2", args{context.Background(), 2 * time.Second, 1 * time.Second}, "time after exceeded"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contextWithDeadline(tt.args.ctx, tt.args.contextDeadline, tt.args.timeAfter); got != tt.want {
				t.Errorf("contextWithDeadline() = %v, want %v", got, tt.want)
			}
		})
	}
}
