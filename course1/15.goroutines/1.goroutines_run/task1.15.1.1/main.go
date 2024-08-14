package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(3 * time.Second)

	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}
}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	out := make(chan string)

	go func() {
		select {
		case <-stop:
			out <- "Горутина завершила работу раньше, чем таймер сработал"
			close(out)
		case <-timer.C:
			out <- "Таймер сработал раньше, чем горутина завершила работу"
			close(out)
		}
	}()
	return out
}
