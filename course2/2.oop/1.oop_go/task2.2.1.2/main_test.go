package main

import (
	"os"
	"testing"
)

var f, _ = os.OpenFile("file.txt", os.O_RDWR, 0777)
var cl = &ConsoleLogger{os.Stdout}

func TestConsoleLogger_Log(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		l       *ConsoleLogger
		args    args
		wantErr bool
	}{
		{"Test console log", cl, args{"Test message"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Log(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("ConsoleLogger.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileLogger_Log(t *testing.T) {
	type args struct {
		message string
	}

	tests := []struct {
		name    string
		l       *FileLogger
		args    args
		wantErr bool
	}{
		{"Test file log", &FileLogger{f}, args{"Test message"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Log(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("FileLogger.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogAll(t *testing.T) {
	type args struct {
		loggers []Logger
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test file log", args{[]Logger{&FileLogger{f}, cl}, "Test message"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogAll(tt.args.loggers, tt.args.message)
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
