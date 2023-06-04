package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 限价频道
// 获取衍生品交易的最高买价和最低卖价。限价有变化时，每5秒推送一次数据，限价没变化时，不推送数据

type HandlerPriceLimit func(EventPriceLimit)

type EventPriceLimit struct {
	Arg  ws.Args      `json:"arg"`
	Data []PriceLimit `json:"data"`
}

type PriceLimit struct {
	InstId  string `json:"instId"`
	BuyLmt  string `json:"buyLmt"`
	SellLmt string `json:"sellLmt"`
	Ts      int64  `json:"ts,string"`
}

// default subscribe
func SubscribePriceLimit(instId string, handler HandlerPriceLimit, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "price-limit",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventPriceLimit
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
