package market

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetCandlesLimitNumPerSec = 20
const GetCandlesLimitRule = "IP"

func NewGetCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}

type GetCandlesParam struct {
	InstId string `url:"instId"`
	Bar    string `url:"bar,omitempty"`
	After  int64  `url:"after,omitempty"`
	Before int64  `url:"before,omitempty"`
	Limit  int    `url:"limit,omitempty"`
}

// 获取交易产品K线数据
// 获取K线数据。K线数据按请求的粒度分组返回，K线数据每个粒度最多可获取最近1,440条。
type GetCandlesResponse struct {
	api.Response
	Data [][]string `json:"data"`
}
