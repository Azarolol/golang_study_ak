package main

import (
	"reflect"
	"testing"
)

var product = &Product{
	ProductID:     ProductCocaCola,
	Sells:         []float64{10, 20, 30},
	Buys:          []float64{5, 15, 25},
	CurrentPrice:  35,
	ProfitPercent: 10,
}

var statProfit = NewStatisticProfit(
	WithAverageProfit,
	WithAverageProfitPercent,
	WithCurrentProfit,
	WithDifferenceProfit,
	WithAllData,
).(*StatisticProfit)

func TestStatisticProfit_GetAllData(t *testing.T) {
	statProfit.SetProduct(product)
	tests := []struct {
		name string
		s    *StatisticProfit
		want []float64
	}{
		{"Test get all data", statProfit, []float64{5, 33.33333333333333, 11.666666666666664, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetAllData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatisticProfit.GetAllData() = %v, want %v", got, tt.want)
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
