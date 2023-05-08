package market

import (
	"go-okx/rest/api"
)

const GetTradesLimitNumPerSec = 50
const GetTradesLimitRule = "IP"

func NewGetTrades(param *GetTradesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/trades",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTradesResponse{}
}

type GetTradesParam struct {
	InstId string `url:"instId"`
	Limit  string `url:"limit"` // 分页返回的结果集数量, 最大为500, 默认100
}

type GetTradesResponse struct {
	api.Response
	Data []TradesData `json:"data"`
}

type TradesData struct {
	InstId  string `json:"instId"`  // 产品ID
	TradeId string `json:"tradeId"` //成交ID
	Px      string `json:"px"`      //成交价格
	Sz      string `json:"sz"`      //成交数量
	Side    string `json:"side"`    // 成交方向: buy; sell
	Ts      string `json:"ts"`      // 成交时间
}
