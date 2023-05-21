package finance

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetActiveOrdersLimitNumPerSec = 3
const GetActiveOrdersLimitRule = "UserID"

func NewGetActiveOrders(param *GetActiveOrdersParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/finance/staking-defi/orders-active",
		Method: api.MethodGet,
		Param:  param,
	}, &GetActiveOrdersResponse{}
}

type GetActiveOrdersParam struct {
	ProductID    string `json:"productId"`    //项目ID
	ProtocolType string `json:"protocolType"` //项目类型
	Ccy          string `json:"ccy"`          //投资币种
	State        string `json:"state"`        //订单状态
	// 8: 待上车（预约中）
	// 13: 订单取消中
	// 9: 上链中
	// 1: 收益中
	// 2: 赎回中
}

type GetActiveOrdersResponse struct {
	api.Response
	Data []struct {
		Ccy          string `json:"ccy"`
		OrdId        string `json:"ordId"`
		ProductId    string `json:"productId"`
		State        string `json:"state"`
		Protocol     string `json:"protocol"`
		ProtocolType string `json:"protocolType"`
		Term         string `json:"term"`
		Apy          string `json:"apy"`
		InvestData   []struct {
			Ccy string `json:"ccy"`
			Amt string `json:"amt"`
		} `json:"investData"`
		EarningData []struct {
			Ccy         string `json:"ccy"`
			EarningType string `json:"earningType"`
			Earnings    string `json:"earnings"`
		} `json:"EarningData"`
		PurchasedTime string `json:"purchasedTime"`
	} `json:"data"`
}
