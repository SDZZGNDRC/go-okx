package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// 资金费率频道
// 获取永续合约资金费率，30秒到90秒内推送一次数据

type HandlerFundingRate func(EventFundingRate)

type EventFundingRate struct {
	Arg  ws.Args       `json:"arg"`
	Data []FundingRate `json:"data"`
}

type FundingRate struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     int64  `json:"fundingTime,string"`
	InstId          string `json:"instId"`
	InstType        string `json:"instType"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime int64  `json:"nextFundingTime,string"`
}

// default subscribe
func SubscribeFundingRate(instId string, handler HandlerFundingRate, handlerError ws.HandlerError, simulated bool) error {
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
