package test

import (
	"errors"
	"reflect"
	"strategy/client"
	"testing"
	"time"
)

type MockIndicatorSMA struct {
	wantErr bool
}

func (i MockIndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	if i.wantErr {
		return nil, errors.New("test error")
	}
	return []float64{4.0, 5.0, 6.0}, nil
}

type MockIndicatorEMA struct {
	wantErr bool
}

func (i MockIndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	if i.wantErr {
		return nil, errors.New("test error")
	}
	return []float64{1.0, 2.0, 3.0}, nil
}

func TestGeneralIndicator_GetData(t *testing.T) {
	type args struct {
		pair      string
		period    int
		from      time.Time
		to        time.Time
		indicator client.Indicatorer
	}
	tests := []struct {
		name    string
		i       client.GeneralIndicator
		args    args
		want    []float64
		wantErr bool
	}{
		{"Test general indicator get data mock ema", client.GeneralIndicator{}, args{"", 0, time.Now(), time.Now(), &MockIndicatorEMA{false}}, []float64{1.0, 2.0, 3.0}, false},
		{"Test general indicator get data mock ema error", client.GeneralIndicator{}, args{"", 0, time.Now(), time.Now(), &MockIndicatorEMA{true}}, nil, true},
		{"Test general indicator get data mock sma", client.GeneralIndicator{}, args{"", 0, time.Now(), time.Now(), &MockIndicatorSMA{false}}, []float64{4.0, 5.0, 6.0}, false},
		{"Test general indicator get data mock sma error", client.GeneralIndicator{}, args{"", 0, time.Now(), time.Now(), &MockIndicatorSMA{true}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetData(tt.args.pair, tt.args.period, tt.args.from, tt.args.to, tt.args.indicator)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneralIndicator.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneralIndicator.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestGeneralIndicator_GetData$ strategy/test

ok  	strategy/test	0.273s
*/
