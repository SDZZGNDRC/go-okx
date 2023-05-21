package statistics

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetLSARatioLimitNumPerSec = 2.5
const GetLSARatioLimitRule = "IP"

func NewGetLSARatio(param *GetLSARatioParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/contracts/long-short-account-ratio",
		Method: api.MethodGet,
		Param:  param,
	}, &GetLSARatioResponse{}
}

type GetLSARatioParam struct {
	Ccy    string `url:"ccy"`
	Begin  string `url:"begin,omitempty"`
	End    string `url:"end,omitempty"`
	Period string `url:"period,omitempty"`
}

type LSARatioData struct {
	Item [2]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetLSARatioResponse struct {
	api.Response
	Data []LSARatioData `json:"data"`
}
