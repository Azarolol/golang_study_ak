package test

import (
	"reflect"
	"testing"
	"time"

	"exmo_mock/client"
	"exmo_mock/mock"

	"go.uber.org/mock/gomock"
)

func Test_GetTicker(t *testing.T) {
	expected := client.Ticker{
		"BTC_USD": {
			BuyPrice:  "589.06",
			SellPrice: "592",
			LastTrade: "591.221",
			High:      "602.082",
			Low:       "584.51011695",
			Avg:       "591.14698808",
			Vol:       "167.59763535",
			VolCurr:   "99095.17162071",
			Updated:   1470250973,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockExchanger(ctrl)

	mock.EXPECT().GetTicker().Return(expected, nil)

	ticker, err := mock.GetTicker()
	if err != nil {
		t.Fatalf("error getting ticker: %v", err)
	}

	if ok := reflect.DeepEqual(expected, ticker); !ok {
		t.Fatalf("wrong result. Want: %v, got: %v", expected, ticker)
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^Test_GetTicker$ exmo_mock/test

ok  	exmo_mock/test	0.288s
*/

func Test_GetCurrencies(t *testing.T) {
	expected := client.Currencies{
		"USD",
		"EUR",
		"BTC",
		"DOGE",
		"LTC",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockExchanger(ctrl)

	mock.EXPECT().GetCurrencies().Return(expected, nil)

	currencies, err := mock.GetCurrencies()
	if err != nil {
		t.Fatalf("error getting currencies: %v", err)
	}

	if ok := reflect.DeepEqual(expected, currencies); !ok {
		t.Fatalf("wrong result. Want: %v, got: %v", expected, currencies)
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^Test_GetCurrencies$ exmo_mock/test

ok  	exmo_mock/test	0.319s
*/

func Test_GetTrades(t *testing.T) {
	expected := client.Trades{
		"BTC_USD": []client.Pair{
			{
				TradeID:  3,
				Type:     "sell",
				Price:    "100",
				Quantity: "1",
				Amount:   "100",
				Date:     1435488248,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockExchanger(ctrl)

	mock.EXPECT().GetTrades("BTC_USD").Return(expected, nil)

	trades, err := mock.GetTrades("BTC_USD")
	if err != nil {
		t.Fatalf("error getting trades: %v", err)
	}

	if ok := reflect.DeepEqual(expected, trades); !ok {
		t.Fatalf("wrong result. Want: %v, got: %v", expected, trades)
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^Test_GetTrades$ exmo_mock/test

ok  	exmo_mock/test	0.275s
*/

func Test_GetCandlesHistory(t *testing.T) {
	expected := client.CandlesHistory{
		Candles: []client.Candle{
			{
				T: 1585557000000,
				O: 6590.6164,
				C: 6602.3624,
				H: 6618.78965693,
				L: 6579.054,
				V: 6.932754980000013,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockExchanger(ctrl)

	start := time.Now()
	end := time.Now().Add(-2 * time.Hour)

	mock.EXPECT().GetCandlesHistory("BTC_USD", 30, start, end).Return(expected, nil)

	trades, err := mock.GetCandlesHistory("BTC_USD", 30, start, end)
	if err != nil {
		t.Fatalf("error getting candles: %v", err)
	}

	if ok := reflect.DeepEqual(expected, trades); !ok {
		t.Fatalf("wrong result. Want: %v, got: %v", expected, trades)
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^Test_GetCandlesHistory$ exmo_mock/test

ok  	exmo_mock/test	0.264s
*/

func Test_OrderBook(t *testing.T) {
	expected := client.OrderBook{
		"BTC_USD": client.OrderBookPair{
			AskQuantity: "3",
			AskAmount:   "500",
			AskTop:      "100",
			BidQuantity: "1",
			BidAmount:   "99",
			BidTop:      "99",
			Ask:         [][]int{{100, 1, 100}, {200, 2, 400}},
			Bid:         [][]int{{99, 1, 99}},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockExchanger(ctrl)

	mock.EXPECT().GetOrderBook(100, "BTC_USD").Return(expected, nil)

	trades, err := mock.GetOrderBook(100, "BTC_USD")
	if err != nil {
		t.Fatalf("error getting order book: %v", err)
	}

	if ok := reflect.DeepEqual(expected, trades); !ok {
		t.Fatalf("wrong result. Want: %v, got: %v", expected, trades)
	}
}

/*
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^Test_OrderBook$ exmo_mock/test

ok  	exmo_mock/test	0.267s
*/
