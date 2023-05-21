package trade

import (
	"github.com/SDZZGNDRC/go-okx/rest/api"
)

const GetOrderLimitNumPerSec = 30

// Special_1:
// 衍生品：UserID + (instrumentType + underlying)
// 币币和币币杠杆：UserID + instrumentID
const GetOrderLimitRule = "Special_1"

func NewGetOrder(param *GetOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

type GetOrderParam struct {
	InstId  string `url:"instId"`
	OrdId   string `url:"ordId,omitempty"`
	ClOrdId string `url:"clOrdId,omitempty"`
}

// 获取订单信息
// 查订单信息
type GetOrderResponse struct {
	api.Response
	Data []OrderDetail `json:"data"`
}
