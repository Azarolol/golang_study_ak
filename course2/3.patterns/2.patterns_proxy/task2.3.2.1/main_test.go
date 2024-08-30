package main

import (
	"reflect"
	"testing"
)

type MockIndicator struct{}

func (MockIndicator) StochPrice() ([]float64, []float64) {
	return []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}
}
func (MockIndicator) RSI(period int) ([]float64, []float64) {
	return []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}
}
func (MockIndicator) StochRSI(rsiPeriod int) ([]float64, []float64) {
	return []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}
}
func (MockIndicator) SMA(period int) []float64 {
	return []float64{1.0, 2.0, 3.0}
}
func (MockIndicator) MACD() ([]float64, []float64) {
	return []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}
}
func (MockIndicator) EMA() []float64 {
	return []float64{1.0, 2.0, 3.0}
}

func TestLinesProxy_StochPrice(t *testing.T) {
	tests := []struct {
		name  string
		tr    *LinesProxy
		want  []float64
		want1 []float64
	}{
		{"Test lines proxy stoch price mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tr.StochPrice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.StochPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LinesProxy.StochPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_RSI(t *testing.T) {
	type args struct {
		period int
	}
	tests := []struct {
		name  string
		tr    *LinesProxy
		args  args
		want  []float64
		want1 []float64
	}{
		{"Test lines proxy rsi mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, args{}, []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tr.RSI(tt.args.period)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.RSI() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LinesProxy.RSI() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_StochRSI(t *testing.T) {
	type args struct {
		rsiPeriod int
	}
	tests := []struct {
		name  string
		tr    *LinesProxy
		args  args
		want  []float64
		want1 []float64
	}{
		{"Test lines proxy stoch rsi mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, args{}, []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tr.StochRSI(tt.args.rsiPeriod)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.StochRSI() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LinesProxy.StochRSI() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_MACD(t *testing.T) {
	tests := []struct {
		name  string
		tr    *LinesProxy
		want  []float64
		want1 []float64
	}{
		{"Test lines proxy macd mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, []float64{1.0, 2.0, 3.0}, []float64{4.0, 5.0, 6.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tr.MACD()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.MACD() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LinesProxy.MACD() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_EMA(t *testing.T) {
	tests := []struct {
		name string
		tr   *LinesProxy
		want []float64
	}{
		{"Test lines proxy ema mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, []float64{1.0, 2.0, 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.EMA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.EMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinesProxy_SMA(t *testing.T) {
	type args struct {
		period int
	}
	tests := []struct {
		name string
		tr   *LinesProxy
		args args
		want []float64
	}{
		{"Test lines proxy ema mock", &LinesProxy{&MockIndicator{}, make(map[string][]float64)}, args{}, []float64{1.0, 2.0, 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.SMA(tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesProxy.SMA() = %v, want %v", got, tt.want)
			}
		})
	}
}
