package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 指数行情频道
// 获取指数的行情数据。每100ms有变化就推送一次数据，否则一分钟推一次。

type HandlerIndexTickers func(EventIndexTickers)

type EventIndexTickers struct {
	Arg  ws.Args        `json:"arg"`
	Data []IndexTickers `json:"data"`
}

type IndexTickers struct {
	InstId  string `json:"instId"`
	IdxPx   string `json:"idxPx"`
	High24h string `json:"high24h"`
	Low24h  string `json:"low24h"`
	Open24h string `json:"open24h"`
	SodUtc0 string `json:"sodUtc0"`
	SodUtc8 string `json:"sodUtc8"`
	Ts      int64  `json:"ts,string"`
}

// default subscribe
func SubscribeIndexTickers(instId string, handler HandlerIndexTickers, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "index-tickers",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventIndexTickers
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
