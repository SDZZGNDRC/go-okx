package public

import (
	"encoding/json"
	"go-okx/ws"
)

// 期权公共成交频道
// 获取最近的期权成交数据，有成交数据就推送，每次推送仅包含一条成交数据。

type HandlerOptDeal func(EventOptDeal)

type EventOptDeal struct {
	Arg  ws.Args   `json:"arg"`
	Data []OptDeal `json:"data"`
}

type OptDeal struct {
	FillVol    string `json:"fillVol,omitempty"`
	FwdPx      string `json:"fwdPx,omitempty"`
	IndexPx    string `json:"indexPx,omitempty"`
	InstFamily string `json:"instFamily,omitempty"`
	InstID     string `json:"instId,omitempty"`
	MarkPx     string `json:"markPx,omitempty"`
	OptType    string `json:"optType,omitempty"`
	Px         string `json:"px,omitempty"`
	Side       string `json:"side,omitempty"`
	Sz         string `json:"sz,omitempty"`
	TradeID    string `json:"tradeId,omitempty"`
	Ts         string `json:"ts,omitempty"`
}

// default subscribe
func SubscribeOptDeal(args *ws.Args, handler HandlerOptDeal, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventOptDeal
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
