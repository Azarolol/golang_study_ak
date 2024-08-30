package main

import (
	"facade/client"
	"fmt"
	"time"
)

func main() {
	exchange := client.NewExmo()
	dashboard := client.NewDashboard(exchange)
	data, err := dashboard.GetDashboard("BTC_USD", client.WithCandlesHistory(30, time.Now().Add(-time.Hour*24*5), time.Now()),
		client.WithIndicatorOpts(
			client.IndicatorOpt{
				Name:      "SMA",
				Periods:   []int{5, 10, 20},
				Indicator: client.NewIndicator(),
			},
			client.IndicatorOpt{
				Name:      "EMA",
				Periods:   []int{5, 10, 20},
				Indicator: client.NewIndicator(),
			},
		),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
