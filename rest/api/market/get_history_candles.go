package market

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetHistoryCandlesLimitNumPerSec = 10
const GetHistoryCandlesLimitRule = "IP"

func NewGetHistoryCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/history-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
