package main

import "testing"

func TestSlowMover_Move(t *testing.T) {
	tests := []struct {
		name string
		m    SlowMover
		want string
	}{
		{"Test slow mover move", SlowMover{BaseMover{10}}, "Slow mover... Moving at speed: 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Move(); got != tt.want {
				t.Errorf("SlowMover.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastMover_Move(t *testing.T) {
	tests := []struct {
		name string
		m    FastMover
		want string
	}{
		{"Test fast mover move", FastMover{BaseMover{100}}, "Fast mover! Moving at speed: 100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Move(); got != tt.want {
				t.Errorf("FastMover.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMover_Move(t *testing.T) {
	tests := []struct {
		name string
		m    BaseMover
		want string
	}{
		{"Test base mover move", BaseMover{50}, "Moving at speed: 50"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Move(); got != tt.want {
				t.Errorf("BaseMover.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMover_Speed(t *testing.T) {
	tests := []struct {
		name string
		m    BaseMover
		want int
	}{
		{"Test base mover speed", BaseMover{50}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Speed(); got != tt.want {
				t.Errorf("BaseMover.Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMover_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		m    BaseMover
		want int
	}{
		{"Test base mover max speed", BaseMover{50}, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MaxSpeed(); got != tt.want {
				t.Errorf("BaseMover.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMover_MinSpeed(t *testing.T) {
	tests := []struct {
		name string
		m    BaseMover
		want int
	}{
		{"Test base mover min speed", BaseMover{50}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MinSpeed(); got != tt.want {
				t.Errorf("BaseMover.MinSpeed() = %v, want %v", got, tt.want)
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
