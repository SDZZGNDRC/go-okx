package trade

import (
	"github.com/SDZZGNDRC/go-okx/rest/api"
)

const GetOrdersPendingLimitNumPerSec = 30
const GetOrdersPendingLimitRule = "UserID"

func NewGetOrdersPending(param *GetOrdersPendingParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrdersPendingResponse{}
}

type GetOrdersPendingParam struct {
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdType  string `url:"ordType,omitempty"`
	State    string `url:"state,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Limit    string `url:"limit,omitempty"`
}

// 获取未成交订单列表
// 获取当前账户下所有未成交订单信息
type GetOrdersPendingResponse struct {
	api.Response
	Data []OrderDetail `json:"data"`
}
