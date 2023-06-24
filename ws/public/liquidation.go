package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 平台公共爆仓单频道
// 获取爆仓单信息。产品ID维度最多一秒推一条爆仓单信息。

type HandlerLiquidation func(EventLiquidation)

type EventLiquidation struct {
	Arg  ws.Args       `json:"arg"`
	Data []Liquidation `json:"data"`
}

type Liquidation struct {
	Details []struct {
		BkLoss  string `json:"bkLoss,omitempty"`
		BkPx    string `json:"bkPx,omitempty"`
		Ccy     string `json:"ccy,omitempty"`
		PosSide string `json:"posSide,omitempty"`
		Side    string `json:"side,omitempty"`
		Sz      string `json:"sz,omitempty"`
		Ts      string `json:"ts,omitempty"`
	} `json:"details,omitempty"`
	InstID   string `json:"instId,omitempty"`
	InstType string `json:"instType,omitempty"`
	Uly      string `json:"uly,omitempty"`
}

// default subscribe
func SubscribeLiquidation(args *ws.Args, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	h := func(message []byte) {
		var event EventLiquidation
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
