package public

import (
	"encoding/json"

	"go-okx/ws"
)

// K线频道
// 获取K线数据，推送频率最快是间隔1秒推送一次数据。
// 频道名:
// candle3M candle1M
// candle1W
// candle1D candle2D candle3D candle5D
// candle12H candle6H candle4H candle2H candle1H
// candle30m candle15m candle5m candle3m candle1m
// candle3Mutc candle1Mutc candle1Wutc candle1Dutc candle2Dutc candle3Dutc candle5Dutc candle12Hutc candle6Hutc

type HandlerKline func(EventKline)

type EventKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeKline(args *ws.Args, handler HandlerKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
