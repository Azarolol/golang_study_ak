package client

import (
	"errors"

	"github.com/cinar/indicator"
)

type Indicatorer interface {
	SMA(period int) ([]float64, error)
	EMA(period int) ([]float64, error)
	LoadCandles(candles CandlesHistory)
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Indicator struct {
	lines *Lines
}

func NewIndicator() *Indicator {
	return &Indicator{}
}

func (i *Indicator) SMA(period int) ([]float64, error) {
	if i.lines.closing == nil {
		return nil, errors.New("no data")
	}
	return CalculateSMA(i.lines.closing, period), nil
}

func (i *Indicator) EMA(period int) ([]float64, error) {
	if i.lines.closing == nil {
		return nil, errors.New("no data")
	}
	return CalculateEMA(i.lines.closing, period), nil
}

func (i *Indicator) LoadCandles(candles CandlesHistory) {
	i.lines = LoadCandles(candles)
}

// Функция для расчет простого скользящего среднего (SMA)
func CalculateSMA(closing []float64, period int) []float64 {
	return indicator.Sma(period, closing)
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func CalculateEMA(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}

func LoadCandles(candles CandlesHistory) *Lines {
	t := &Lines{}
	for i := 0; i < len(candles.Candles); i++ {
		t.closing = append(t.closing, candles.Candles[i].C)
		t.low = append(t.low, candles.Candles[i].L)
		t.high = append(t.high, candles.Candles[i].H)
	}

	return t
}
