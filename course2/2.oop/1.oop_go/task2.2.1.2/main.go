package main

import (
	"log"
	"os"
)

type Logger interface {
	Log(messge string) error
}

type ConsoleLogger struct {
	Writer *os.File
}

type FileLogger struct {
	File *os.File
}

type RemoteLogger struct {
	Address string
}

func (l *ConsoleLogger) Log(message string) error {
	_, err := l.Writer.WriteString(message)
	return err
}

func (l *FileLogger) Log(message string) error {
	_, err := l.File.WriteString(message)
	return err
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	f, err := os.OpenFile("file.txt", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileLogger := &FileLogger{File: f}

	loggers := []Logger{consoleLogger, fileLogger}
	LogAll(loggers, "This is a test log message.")
}
