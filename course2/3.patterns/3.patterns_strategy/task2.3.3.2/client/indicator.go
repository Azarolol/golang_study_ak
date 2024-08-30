package client

import (
	"time"

	"github.com/cinar/indicator"
)

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}

type GeneralIndicator struct{}

func (i GeneralIndicator) GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, 3, period, from, to)
}

type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type IndicatorSMA struct {
	exchange Exchanger
}

func NewIndicatorSMA(exchange Exchanger) Indicatorer {
	return &IndicatorSMA{
		exchange: exchange,
	}
}

func (i IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	closing, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}

	return CalculateSMA(closing, period), nil
}

type IndicatorEMA struct {
	exchange Exchanger
}

func NewIndicatorEMA(exchange Exchanger) Indicatorer {
	return &IndicatorEMA{
		exchange: exchange,
	}
}

func (i IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	closing, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}

	return CalculateEMA(closing, period), nil
}

// Функция для расчет простого скользящего среднего (SMA)
func CalculateSMA(closing []float64, period int) []float64 {
	return indicator.Sma(period, closing)
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func CalculateEMA(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}
