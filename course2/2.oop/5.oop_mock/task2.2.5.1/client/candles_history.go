package client

import "encoding/json"

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

func UnmarshalCandlesHistory(data []byte) (CandlesHistory, error) {
	var ch CandlesHistory
	err := json.Unmarshal(data, &ch)
	return ch, err
}

func (ch *CandlesHistory) Marshal() ([]byte, error) {
	return json.Marshal(ch)
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}
