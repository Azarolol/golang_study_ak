package client

import (
	"bytes"
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

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) Exchanger {
	exmo := &Exmo{
		http.DefaultClient,
		"https://api.exmo.me/v1",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
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

	return UnmarshalTicker(result)
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	result, err := e.SendRequest(currency, map[string]string{})
	if err != nil {
		return nil, err
	}

	return UnmarshalCurrencies(result)
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	result, err := e.SendRequest(trades, map[string]string{"pair": strings.Join(pairs, ",")})
	if err != nil {
		return nil, err
	}

	return UnmarshalTrades(result)
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	result, err := e.SendRequest(candlesHistory, map[string]string{"symbol": pair, "resolution": strconv.Itoa(limit), "from": fmt.Sprintf("%d", start.UnixNano()), "to": fmt.Sprintf("%d", end.UnixNano())})
	if err != nil {
		return CandlesHistory{}, err
	}

	return UnmarshalCandlesHistory(result)
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	result, err := e.SendRequest(orderBook, map[string]string{"pair": strings.Join(pairs, ","), "limit": strconv.Itoa(limit)})
	if err != nil {
		return nil, err
	}

	return UnmarshalOrderBook(result)
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
