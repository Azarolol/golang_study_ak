package main

import (
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

// FileLogger struct
type FileLogger struct {
	file *os.File
}

func (l FileLogger) Log(message string) {
	l.file.WriteString(message)
}

type LogSystem struct {
	Logger
}

func NewLogSystem(lops ...LogOption) *LogSystem {
	logSystem := &LogSystem{}
	for _, lop := range lops {
		lop(logSystem)
	}
	return logSystem
}

func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.Logger = logger
	}
}

// LogOption functional option type
type LogOption func(*LogSystem)

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}
