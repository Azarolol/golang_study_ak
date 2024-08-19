package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Currencies []string

type OrderBook map[string]OrderBookPair

type Ticker map[string]TickerValue

type Trades map[string][]Pair

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Pair struct {
	TradeID  int    `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int64  `json:"date"`
}

type OrderBookPair struct {
	AskQuantity string  `json:"ask_quantity"`
	AskAmount   string  `json:"ask_amount"`
	AskTop      string  `json:"ask_top"`
	BidQuantity string  `json:"big_quantity"`
	BidAmount   string  `json:"big_amount"`
	BidTop      string  `json:"bid_top"`
	Ask         [][]int `json:"ask"`
	Bid         [][]int `json:"bid"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func (e *Exmo) GetTicker() (Ticker, error) {
	result, err := e.SendRequest(ticker, map[string]string{})
	if err != nil {
		return nil, err
	}

	ticker := make(map[string]TickerValue)

	err = json.Unmarshal(result, &ticker)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	result, err := e.SendRequest(currency, map[string]string{})
	if err != nil {
		return nil, err
	}
	currencies := make([]string, 0)

	err = json.Unmarshal(result, &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	result, err := e.SendRequest(trades, map[string]string{"pair": strings.Join(pairs, ",")})
	if err != nil {
		return nil, err
	}
	trades := make(map[string][]Pair)

	err = json.Unmarshal(result, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	result, err := e.SendRequest(candlesHistory, map[string]string{"symbol": pair, "resolution": strconv.Itoa(limit), "from": fmt.Sprintf("%d", start.UnixNano()), "to": fmt.Sprintf("%d", end.UnixNano())})
	if err != nil {
		return CandlesHistory{}, err
	}
	candlesHistory := CandlesHistory{make([]Candle, 0)}
	err = json.Unmarshal(result, &candlesHistory)
	if err != nil {
		return CandlesHistory{}, err
	}
	return candlesHistory, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	result, err := e.SendRequest(orderBook, map[string]string{"pair": strings.Join(pairs, ","), "limit": strconv.Itoa(limit)})
	if err != nil {
		return nil, err
	}
	orderBook := make(map[string]OrderBookPair)
	err = json.Unmarshal(result, &orderBook)
	if err != nil {
		return nil, err
	}

	return orderBook, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}
	prices := make([]float64, 0)
	for _, candle := range candles.Candles {
		prices = append(prices, candle.C)
	}
	return prices, nil
}

func (e *Exmo) SendRequest(method string, params map[string]string) ([]byte, error) {
	post_params := url.Values{}
	post_params.Add("nonce", nonce())
	for key, value := range params {
		post_params.Add(key, value)
	}

	post_content := post_params.Encode()

	req, _ := http.NewRequest("POST", e.url+method, bytes.NewBuffer([]byte(post_content)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(post_content)))

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return nil, errors.New("http status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func nonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{
		http.DefaultClient,
		"https://api.exmo.me/v1",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

func main() {
	var exchange Exchanger = NewExmo()

	ticker, err := exchange.GetTicker()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ticker:", ticker)

	currencies, err := exchange.GetCurrencies()
	if err != nil {
		panic(err)
	}
	fmt.Println("Currencies:", currencies)

	candlesHistory, err := exchange.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("Candles history:", candlesHistory)

	trades, err := exchange.GetTrades("BTC_USD")
	if err != nil {
		panic(err)
	}
	fmt.Println("Trades:", trades)

	orderBook, err := exchange.GetOrderBook(100, "BTC_USD")
	if err != nil {
		panic(err)
	}
	fmt.Println("Order book:", orderBook)

	closePrice, err := exchange.GetClosePrice("BTC_USD", 30, time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("Close price:", closePrice)
}
