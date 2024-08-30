package main

import (
	"fmt"
	"log"
	"strategy/client"
	"time"
)

func main() {
	var exchange client.Exchanger = client.NewExmo()
	indicatorSMA := client.NewIndicatorSMA(exchange)
	generalIndicator := &client.GeneralIndicator{}
	sma, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sma)

	indicatorEMA := client.NewIndicatorEMA(exchange)
	ema, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ema)
}
