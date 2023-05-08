package market

import "go-okx/rest/api"

const GetBooksLimitNumPerSec = 10
const GetBooksLimitRule = "IP"

func NewGetBooks(param *GetBooksParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/books",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBooksResponse{}
}

type GetBooksParam struct {
	InstId string `url:"instId"`
	Sz     string `url:"sz,omitempty"`
}

type BooksData struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	Ts   int64      `json:"ts,string"`
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetBooksResponse struct {
	api.Response
	Data []BooksData `json:"data"`
}
