package market

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetTickersLimitNumPerSec = 10
const GetTickersLimitRule = "IP"

func NewGetTickers(param *GetTickersParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/tickers",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickersResponse{}
}

type GetTickersParam struct {
	InstType string `url:"instType"`
	Uly      string `url:"uly,omitempty"`
}

// 获取所有产品行情信息
// 获取产品行情信息
type GetTickersResponse struct {
	api.Response
	Data []Ticker `json:"data"`
}
