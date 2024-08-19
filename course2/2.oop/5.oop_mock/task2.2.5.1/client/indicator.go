package client

import (
	"time"

	"github.com/cinar/indicator"
)

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOption func(*Indicator)

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) Indicatorer {
	indicator := &Indicator{exchange: exchange}
	for _, opt := range opts {
		opt(indicator)
	}
	return indicator
}

// Функция для расчет простого скользящего среднего (SMA)
func CalculateSMA(closing []float64, period int) []float64 {
	return indicator.Sma(period, closing)
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func CalculateEMA(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}

func WithCalculateSMA(sma func(data []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateSMA = sma
	}
}

func WithCalculateEMA(ema func(data []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateEMA = ema
	}
}

func (i *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	closePrice, err := i.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	return CalculateSMA(closePrice, period), nil
}

func (i *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	closePrice, err := i.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	return CalculateEMA(closePrice, period), nil
}
