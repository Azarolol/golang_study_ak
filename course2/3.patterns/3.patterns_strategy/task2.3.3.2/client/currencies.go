package client

import "encoding/json"

type Currencies []string

func UnmarshalCurrencies(data []byte) (Currencies, error) {
	var c Currencies
	err := json.Unmarshal(data, &c)
	return c, err
}

func (c *Currencies) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
