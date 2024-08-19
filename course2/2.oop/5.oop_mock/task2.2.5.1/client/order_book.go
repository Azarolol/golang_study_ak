package client

import "encoding/json"

type OrderBook map[string]OrderBookPair

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

func UnmarshalOrderBook(data []byte) (OrderBook, error) {
	var r OrderBook
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OrderBook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
