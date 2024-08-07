package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J")
		t := time.Now()
		fmt.Printf("Текущее время: %02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
		fmt.Printf("Текущая дата: %d-%02d-%02d\n", t.Year(), t.Month(), t.Day())
	}
}
