package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 交易频道
// 获取最近的成交数据，有成交数据就推送，每次推送仅包含一条成交数据。

type HandlerTradesAll func(EventTrades)

type EventTradesAll struct {
	Arg  ws.Args `json:"arg"`
	Data []Trade `json:"data"`
}

type TradeAll struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      int64  `json:"ts,string"`
}

// default subscribe
func SubscribeTradesAll(instId string, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "trades-all",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventTradesAll
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
