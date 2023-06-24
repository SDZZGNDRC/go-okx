package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 行情频道
// 获取产品的最新成交价、买一价、卖一价和24小时交易量等信息。
// 最快100ms推送一次，没有触发事件时不推送，触发推送的事件有：成交、买一卖一发生变动。

type HandlerTickers func(EventTickers)

type EventTickers struct {
	Arg  ws.Args  `json:"arg"`
	Data []Ticker `json:"data"`
}

type Ticker struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        int64  `json:"ts,string"`
}

// default subscribe
func SubscribeTickers(instId string, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "tickers",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventTickers
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
