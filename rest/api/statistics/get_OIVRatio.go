package statistics

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetOIVRatioLimitNumPerSec = 2.5
const GetOIVRatioLimitRule = "IP"

func NewGetOIVRatio(param *GetOIVRatioParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/option/open-interest-volume-ratio",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOIVRatioResponse{}
}

type GetOIVRatioParam struct {
	Ccy    string `url:"ccy"`
	Period string `url:"period,omitempty"`
}

type OIVRatioData struct {
	Item [3]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetOIVRatioResponse struct {
	api.Response
	Data []OIVRatioData `json:"data"`
}
