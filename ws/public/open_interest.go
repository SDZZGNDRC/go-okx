package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 持仓总量频道
// 获取持仓总量，每3s有数据更新推送一次数据

type HandlerOpenInterest func(EventOpenInterest)

type EventOpenInterest struct {
	Arg  ws.Args        `json:"arg"`
	Data []OpenInterest `json:"data"`
}

type OpenInterest struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	Oi       string `json:"oi"`
	OiCcy    string `json:"oiCcy"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeOpenInterest(instId string, handler HandlerOpenInterest, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "open-interest",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventOpenInterest
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
