package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// 预估交割/行权价格频道
// 获取永续合约，交割合约和期权预估交割/行权价。
// 交割/行权预估价只有交割/行权前一小时开始推送预估交割/行权价，有价格变化就推送

type HandlerEstimatedPrice func(EventEstimatedPrice)

type EventEstimatedPrice struct {
	Arg  ws.Args          `json:"arg"`
	Data []EstimatedPrice `json:"data"`
}

type EstimatedPrice struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	SettlePx string `json:"settlePx"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeEstimatedPrice(args *ws.Args, handler HandlerEstimatedPrice, handlerError ws.HandlerError, simulated bool) error {
	args.Channel = "estimated-price"

	h := func(message []byte) {
		var event EventEstimatedPrice
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
