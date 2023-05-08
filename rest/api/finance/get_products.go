package finance

import "go-okx/rest/api"

const GetProductsLimitNumPerSec = 3
const GetProductsLimitRule = "UserID"

func NewGetProducts(param *GetProductsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/finance/staking-defi/offers",
		Method: api.MethodGet,
		Param:  param,
	}, &GetProductsResponse{}
}

type GetProductsParam struct {
	ProductId    string `url:"productId,omitempty"`
	ProtocolType string `url:"protocolType,omitempty"`
	Ccy          string `url:"ccy,omitempty"`
}

type GetProductsResponse struct {
	api.Response
	Data []ProductsData `json:"data"`
}

type ProductsData struct {
	Ccy          string            `json:"ccy"`
	ProductId    string            `json:"productId"`
	Protocol     string            `json:"protocol"`     //项目名称
	ProtocolType string            `json:"protocolType"` //项目类型
	Term         string            `json:"term"`         //活动期限; 活期为0，其他则显示定期天数
	Apy          string            `json:"apy"`          //预估年华; 如果年化为7% ，则该字段为0.07
	EarlyRedeem  bool              `json:"earlyRedeem"`  //项目是否支持提前赎回
	InvestData   []InvestDataItem  `json:"investData"`   //目前用户可用来投资的目标币种信息
	EarningData  []EarningDataItem `json:"earningData"`  //收益信息
	State        string            `json:"state"`        // 项目状态
}
