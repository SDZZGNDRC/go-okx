package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// 标记价格频道
// 获取标记价格，标记价格有变化时，每200ms推送一次数据，标记价格没变化时，每10s推送一次数据

type HandlerMarkPrice func(EventMarkPrice)

type EventMarkPrice struct {
	Arg  ws.Args     `json:"arg"`
	Data []MarkPrice `json:"data"`
}

type MarkPrice struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	MarkPx   string `json:"markPx"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeMarkPrice(instId string, handler HandlerMarkPrice, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel: "mark-price",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventMarkPrice
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
