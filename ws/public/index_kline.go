package public

import (
	"encoding/json"
	"time"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 指数K线频道
// 获取指数的K线数据，推送频率最快是间隔1秒推送一次数据。

type HandlerIndexKline func(EventIndexKline)

type EventIndexKline struct {
	Arg     ws.Args    `json:"arg"`
	Data    [][]string `json:"data"`
	LocalTs int64      `json:"localTs"`
}

// default subscribe
func SubscribeIndexKline(args *ws.Args, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	h := func(message []byte) {
		var event EventIndexKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		event.LocalTs = time.Now().UnixMilli()
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
