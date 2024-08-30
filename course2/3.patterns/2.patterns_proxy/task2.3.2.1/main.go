package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cinar/indicator"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *KLines) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type KLines struct {
	Pair    string   `json:"pair"`
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(t.high, t.low, t.closing)

	return k, d
}

func (t *LinesProxy) StochPrice() ([]float64, []float64) {
	key1 := "k_stochprice"
	key2 := "d_stochprice"
	if v1, ok := t.cache[key1]; ok {
		v2 := t.cache[key2]
		return v1, v2
	}
	v1, v2 := t.lines.StochPrice()
	t.cache[key1] = v1
	t.cache[key2] = v2
	return v1, v2
}

func (t *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, t.closing)

	return rs, rsi
}

func (t *LinesProxy) RSI(period int) ([]float64, []float64) {
	key1 := fmt.Sprintf("rs_%v", period)
	key2 := fmt.Sprintf("rsi_%v", period)
	if v1, ok := t.cache[key1]; ok {
		v2 := t.cache[key2]
		return v1, v2
	}
	v1, v2 := t.lines.RSI(period)
	t.cache[key1] = v1
	t.cache[key2] = v2
	return v1, v2
}

func (t *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := t.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)

	return k, d
}

func (t *LinesProxy) StochRSI(rsiPeriod int) ([]float64, []float64) {
	key1 := fmt.Sprintf("k_stochrsi_%v", rsiPeriod)
	key2 := fmt.Sprintf("d_stochrsi_%v", rsiPeriod)
	if v1, ok := t.cache[key1]; ok {
		v2 := t.cache[key2]
		return v1, v2
	}
	v1, v2 := t.lines.StochRSI(rsiPeriod)
	t.cache[key1] = v1
	t.cache[key2] = v2
	return v1, v2
}

func (t *Lines) MACD() ([]float64, []float64) {
	return indicator.Macd(t.closing)
}

func (t *LinesProxy) MACD() ([]float64, []float64) {
	key1 := "macd"
	key2 := "signal"
	if v1, ok := t.cache[key1]; ok {
		v2 := t.cache[key2]
		return v1, v2
	}
	v1, v2 := t.lines.MACD()
	t.cache[key1] = v1
	t.cache[key2] = v2
	return v1, v2
}

func (t *Lines) EMA() []float64 {
	return indicator.Ema(5, t.closing)
}

func (t *LinesProxy) EMA() []float64 {
	key := "ema"
	if v, ok := t.cache[key]; ok {
		return v
	}
	v := t.lines.EMA()
	t.cache[key] = v
	return v
}

func (t *Lines) SMA(period int) []float64 {
	return indicator.Sma(period, t.closing)
}

func (t *LinesProxy) SMA(period int) []float64 {
	key := fmt.Sprintf("sma_%v", period)
	if v, ok := t.cache[key]; ok {
		return v
	}
	v := t.lines.SMA(period)
	t.cache[key] = v
	return v
}

type LinesProxy struct {
	lines Indicator
	cache map[string][]float64
}

func LoadKlinesProxy(data []byte) *LinesProxy {
	linesProxy := LinesProxy{}
	linesProxy.lines = LoadKlines(data)
	linesProxy.cache = make(map[string][]float64)
	return &linesProxy
}

func LoadKlines(data []byte) *Lines {
	klines, err := UnmarshalKLines(data)
	if err != nil {
		log.Fatal(err)
	}

	t := &Lines{}
	for _, v := range klines.Candles {
		t.closing = append(t.closing, v.C)
		t.low = append(t.low, v.L)
		t.high = append(t.high, v.H)
	}

	return t
}

func LoadCandles(pair string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.exmo.me/v1.1/candles_history?symbol=%s&resolution=30&from=1703056979&to=1705476839", pair), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	pair := "BTC_USD"
	candles := LoadCandles(pair)
	lines := LoadKlinesProxy(candles)
	fmt.Println(lines.RSI(3))
}
