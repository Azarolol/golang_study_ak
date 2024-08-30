package test

import (
	"facade/client"
	"reflect"
	"testing"
)

type MockIndicator struct {
}

func NewMockIndicator() *MockIndicator {
	return &MockIndicator{}
}

func (MockIndicator) SMA(period int) ([]float64, error) {
	return []float64{1.0, 2.0, 3.0}, nil
}

func (MockIndicator) EMA(period int) ([]float64, error) {
	return []float64{4.0, 5.0, 6.0}, nil
}

func (MockIndicator) LoadCandles(candles client.CandlesHistory) {}

func TestDashboard_GetDashboard(t *testing.T) {
	type args struct {
		pair string
		opts []func(*client.Dashboard)
	}
	tests := []struct {
		name    string
		d       *client.Dashboard
		args    args
		want    client.DashboardData
		wantErr bool
	}{
		{"Test get dashboard mock", &client.Dashboard{}, args{"", []func(*client.Dashboard){
			client.WithIndicatorOpts(
				client.IndicatorOpt{
					Name:      "SMA",
					Periods:   []int{5, 10, 20},
					Indicator: NewMockIndicator(),
				},
				client.IndicatorOpt{
					Name:      "EMA",
					Periods:   []int{5, 10, 20},
					Indicator: NewMockIndicator(),
				},
			)}}, client.DashboardData{
			Name:           "Dashboard for pair ",
			CandlesHistory: client.CandlesHistory{},
			Indicators: map[string][]client.IndicatorData{
				"SMA": {
					{Name: "SMA", Period: 5, Indicate: []float64{1.0, 2.0, 3.0}},
					{Name: "SMA", Period: 10, Indicate: []float64{1.0, 2.0, 3.0}},
					{Name: "SMA", Period: 20, Indicate: []float64{1.0, 2.0, 3.0}},
				},
				"EMA": {
					{Name: "EMA", Period: 5, Indicate: []float64{4.0, 5.0, 6.0}},
					{Name: "EMA", Period: 10, Indicate: []float64{4.0, 5.0, 6.0}},
					{Name: "EMA", Period: 20, Indicate: []float64{4.0, 5.0, 6.0}},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetDashboard(tt.args.pair, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dashboard.GetDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dashboard.GetDashboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestDashboard_GetDashboard$ facade/test

ok  	facade/test	0.239s
*/
