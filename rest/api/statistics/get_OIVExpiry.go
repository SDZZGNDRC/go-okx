package statistics

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetOIVExpiryLimitNumPerSec = 2.5
const GetOIVExpiryLimitRule = "IP"

func NewGetOIVExpiry(param *GetOIVExpiryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/option/open-interest-volume-expiry",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOIVExpiryResponse{}
}

type GetOIVExpiryParam struct {
	Ccy    string `url:"ccy"`
	Period string `url:"period,omitempty"`
}

type OIVExpiryData struct {
	Item [6]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetOIVExpiryResponse struct {
	api.Response
	Data []OIVExpiryData `json:"data"`
}
