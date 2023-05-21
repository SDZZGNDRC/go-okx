package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// 指数K线频道
// 获取指数的K线数据，推送频率最快是间隔1秒推送一次数据。

type HandlerIndexKline func(EventIndexKline)

type EventIndexKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeIndexKline(args *ws.Args, handler HandlerIndexKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventIndexKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
