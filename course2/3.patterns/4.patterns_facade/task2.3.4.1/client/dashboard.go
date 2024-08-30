package client

import (
	"errors"
	"time"
)

// Dashboarder must return candles history and indicators with several periods, presets by opts
type Dashboarder interface {
	GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	period         int
	from           time.Time
	to             time.Time
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	period             int
	from               time.Time
	to                 time.Time
}

func NewDashboard(exchange Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func WithCandlesHistory(period int, from, to time.Time) func(*Dashboard) {
	return func(d *Dashboard) {
		d.period = period
		d.from = from
		d.to = to
		d.withCandlesHistory = true
	}
}

func WithIndicatorOpts(opts ...IndicatorOpt) func(*Dashboard) {
	return func(d *Dashboard) {
		d.IndicatorOpts = append(d.IndicatorOpts, opts...)
	}
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	for _, opt := range opts {
		opt(d)
	}

	var candlesHistory CandlesHistory
	var err error

	if d.withCandlesHistory {
		candlesHistory, err = d.exchange.GetCandlesHistory(pair, d.period, d.from, d.to)
		if err != nil {
			return DashboardData{}, errors.New("cant get candles history")
		}
	}

	indicators := make(map[string][]IndicatorData)

	for _, opt := range d.IndicatorOpts {
		datas := make([]IndicatorData, 0)
		opt.Indicator.LoadCandles(candlesHistory)
		switch opt.Name {
		case "SMA":
			for _, period := range opt.Periods {
				value, err := opt.Indicator.SMA(period)
				if err != nil {
					return DashboardData{}, err
				}

				datas = append(datas, IndicatorData{
					Name:     opt.Name,
					Period:   period,
					Indicate: value,
				})
			}
		case "EMA":
			for _, period := range opt.Periods {
				value, err := opt.Indicator.EMA(period)
				if err != nil {
					return DashboardData{}, err
				}

				datas = append(datas, IndicatorData{
					Name:     opt.Name,
					Period:   period,
					Indicate: value,
				})
			}
		default:
			return DashboardData{}, errors.New("unknown indicator name")
		}
		indicators[opt.Name] = datas
	}

	return DashboardData{
		Name:           "Dashboard for pair " + pair,
		CandlesHistory: candlesHistory,
		Indicators:     indicators,
		period:         d.period,
		from:           d.from,
		to:             d.to,
	}, nil
}
