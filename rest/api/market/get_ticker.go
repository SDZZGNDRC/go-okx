package market

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetTickerLimitNumPerSec = 10
const GetTickerLimitRule = "IP"

func NewGetTicker(param *GetTickerParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/ticker",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickerResponse{}
}

type GetTickerParam struct {
	InstId string `url:"instId"`
}

// 获取单个产品行情信息
// 获取产品行情信息
type GetTickerResponse struct {
	api.Response
	Data []Ticker `json:"data"`
}
