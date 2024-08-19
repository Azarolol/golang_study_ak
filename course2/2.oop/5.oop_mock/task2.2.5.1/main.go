package main

import (
	"exmo_mock/client"
	"fmt"
	"time"
)

func main() {
	var exchange client.Exchanger = client.NewExmo()
	ind := client.NewIndicator(exchange, client.WithCalculateEMA(client.CalculateEMA), client.WithCalculateSMA(client.CalculateSMA))
	sma, err := ind.SMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sma)
	ema, err := ind.EMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ema)
}
