package statistics

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetOIVStrikeLimitNumPerSec = 2.5
const GetOIVStrikeLimitRule = "IP"

func NewGetOIVStrike(param *GetOIVStrikeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/rubik/stat/option/open-interest-volume-strike",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOIVStrikeResponse{}
}

type GetOIVStrikeParam struct {
	Ccy     string `url:"ccy"`
	ExpTime string `url:"expTime"`
	Period  string `url:"period,omitempty"`
}

type OIVStrikeData struct {
	Item [6]string
}

// 获取产品深度
// 获取产品深度列表(订单列表)
type GetOIVStrikeResponse struct {
	api.Response
	Data []OIVStrikeData `json:"data"`
}
