package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 资金费率频道
// 获取永续合约资金费率，30秒到90秒内推送一次数据

type HandlerFundingRate func(EventFundingRate)

type EventFundingRate struct {
	Arg  ws.Args       `json:"arg"`
	Data []FundingRate `json:"data"`
}

type FundingRate struct {
	InstType        string `json:"instType"`
	InstId          string `json:"instId"`
	Method          string `json:"method"`
	FundingRate     string `json:"fundingRate"`
	NextFundingRate string `json:"nextFundingRate"`
	FundingTime     int64  `json:"fundingTime,string"`
	NextFundingTime int64  `json:"nextFundingTime,string"`
	MinFundingRate  string `json:"minFundingRate"`
	MaxFundingRate  string `json:"maxFundingRate"`
	SettState       string `json:"settState"`
	Premium         string `json:"premium"`
	Ts              int64  `json:"ts,string"`
}

// default subscribe
func SubscribeFundingRate(instId string, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	args := &ws.Args{
		Channel: "funding-rate",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventFundingRate
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
