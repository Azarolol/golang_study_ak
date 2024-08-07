package main

import (
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)

	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	println("averageTime", averageTime/time.Second)
	println("sinceProgramStarted", sinceProgramStarted/time.Second)
	println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)
	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	var t time.Time
	for i := 0; i < count; i++ {
		wg.Add(1)
		limit <- struct{}{}
		t = time.Now()
		go func(id int, t time.Time) {
			OrderProcessing(id, t, limit)
		}(i+1, t)
	}
}

func OrderProcessing(id int, t time.Time, limit chan struct{}) {
	if _, ok := completeOrders[id]; ok {
		orders[id-1].Complete = true
	}
	time.Sleep(1 * time.Second)
	processTimes <- time.Since(t)
	wg.Done()
	<-limit
}

func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{i + 1, false}
	}
	return orders
}

func GenerateCompleteOrders(maxOrderID int) map[int]bool {
	rand.Seed(time.Now().UnixNano())
	completeOrders := make(map[int]bool)
	for i := 1; i <= maxOrderID; i++ {
		if rand.Intn(2) == 1 {
			completeOrders[i] = true
		}
	}
	return completeOrders
}
