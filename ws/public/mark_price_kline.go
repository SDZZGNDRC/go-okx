package public

import (
	"encoding/json"

	"go-okx/ws"
)

// 标记价格K线频道
// 获取标记价格的K线数据，推送频率最快是间隔1秒推送一次数据。
// 频道名:
// mark-price-candle3M
// mark-price-candle1M
// mark-price-candle1W
// mark-price-candle1D
// mark-price-candle2D
// mark-price-candle3D
// mark-price-candle5D
// mark-price-candle12H
// mark-price-candle6H
// mark-price-candle4H
// mark-price-candle2H
// mark-price-candle1H
// mark-price-candle30m
// mark-price-candle15m
// mark-price-candle5m
// mark-price-candle3m
// mark-price-candle1m
// mark-price-candle3Mutc
// mark-price-candle1Mutc
// mark-price-candle1Wutc
// mark-price-candle1Dutc
// mark-price-candle2Dutc
// mark-price-candle3Dutc
// mark-price-candle5Dutc
// mark-price-candle12Hutc
// mark-price-candle6Hutc

type HandlerMarkPriceKline func(EventMarkPriceKline)

type EventMarkPriceKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeMarkPriceKline(args *ws.Args, handler HandlerMarkPriceKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventMarkPriceKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
